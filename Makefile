APPNAME := lambda-golang-containers
BRANCH ?= master
STAGE ?= dev

GIT_HASH := $(shell git rev-parse --short HEAD)
ARCH ?= arm64
PLATFORM := linux/$(ARCH)

.PHONY: deploy-repository
deploy-repository:
	@echo "--- deploy stack $(APPNAME)-$(STAGE)-$(BRANCH)-repository"

	@sam deploy \
		--no-fail-on-empty-changeset \
		--template-file sam/app/repository.yaml \
		--capabilities CAPABILITY_IAM \
		--tags "environment=$(STAGE)" "branch=$(BRANCH)" "service=$(APPNAME)" \
		--stack-name $(APPNAME)-$(STAGE)-$(BRANCH)-repository \
		--parameter-overrides AppName=$(APPNAME) Stage=$(STAGE) Branch=$(BRANCH)

.PHONY: docker-login
docker-login:
	$(eval KO_DOCKER_REPO := $(shell aws ssm get-parameter --name '/config/$(STAGE)/$(BRANCH)/$(APPNAME)/repository_uri' --query 'Parameter.Value' --output text))

	aws ecr get-login-password | docker login --username AWS --password-stdin $(KO_DOCKER_REPO)

.PHONY: docker-publish-local
docker-publish-local:
	ko publish --local --platform=$(PLATFORM) --image-label arch=$(ARCH) --image-label git_hash=$(GIT_HASH) --bare ./cmd/api-lambda

.PHONY: deploy-api
deploy-api:
	@echo "--- deploy stack $(APPNAME)-$(STAGE)-$(BRANCH)-api"
	$(eval KO_DOCKER_REPO := $(shell aws ssm get-parameter --name '/config/$(STAGE)/$(BRANCH)/$(APPNAME)/repository_uri' --query 'Parameter.Value' --output text))
	$(eval IMAGE_URL := $(shell KO_DOCKER_REPO=$(KO_DOCKER_REPO) ko publish --platform=$(PLATFORM) --image-label arch=$(ARCH) --image-label git_hash=$(GIT_HASH) --bare ./cmd/api-lambda))

	sam deploy \
		--no-fail-on-empty-changeset \
		--s3-bucket $(SAM_BUCKET) \
		--s3-prefix sam/$(GIT_HASH) \
		--template-file sam/app/api.yaml \
		--image-repository $(KO_DOCKER_REPO) \
		--capabilities CAPABILITY_IAM \
		--tags "environment=$(STAGE)" "branch=$(BRANCH)" "service=$(APPNAME)" \
		--stack-name $(APPNAME)-$(STAGE)-$(BRANCH)-api \
		--parameter-overrides AppName=$(APPNAME) Stage=$(STAGE) Branch=$(BRANCH) ImageUri=$(IMAGE_URL) LambdaArchitecture=$(ARCH)