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

.PHONY: publish-images
publish-images:
	ko publish --platform=$(PLATFORM) --image-label arch=$(ARCH) --image-label git_hash=$(GIT_HASH) --bare ./cmd/api-lambda

.PHONY: deploy-api
deploy-api:
	@echo "--- deploy stack $(APPNAME)-$(STAGE)-$(BRANCH)-api"
	$(eval IMAGE_URL := $(shell ko publish --platform=$(PLATFORM) --image-label arch=$(ARCH) --image-label git_hash=$(GIT_HASH) --bare ./cmd/api-lambda))

	@sam deploy \
		--no-fail-on-empty-changeset \
		--template-file sam/app/api.yaml \
		--capabilities CAPABILITY_IAM \
		--tags "environment=$(STAGE)" "branch=$(BRANCH)" "service=$(APPNAME)" \
		--stack-name $(APPNAME)-$(STAGE)-$(BRANCH)-api \
		--parameter-overrides AppName=$(APPNAME) Stage=$(STAGE) Branch=$(BRANCH) ImageUri=$(IMAGE_URL)