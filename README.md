# lambda-golang-containers

This project illustrates how to use docker containers for golang based lambdas with support for arm64 and x86_64 architectures.

# overview

This project doesn't require docker to be installed as it uses [ko](https://github.com/google/ko) to publish the containers.

1. You can build and publish containers without having docker installed.
2. Building on Go cross compiling, you can easily build binaries for arm64 or amd64 without complex tool chains.
3. Containers built with this tool are under 10MB, so they are easy to manage and don't cost much in ECR.
4. When pushed the containers update the latest tag, so you only need to keep a couple of untagged containers in case of rollback of cloudformation.
5. It is easy to switch between amd64 and arm64 builds, and lambda runtimes purely by changing a flag in the `Makefile`.

**Note:** Although ko tags the image under latest, the lambda is deployed using the more specific sha of the container so it won't break when there is a rollback as long as you keep a couple of untagged versions.

# Using ko

The `ko` tool built by google makes it really easy to use distroless containers without docker installed.

Although the readme for `ko` provides an example of pushing Go based lambdas I found this didn't work without using:

* The `debug` version of the distroless containers which includes busybox for executing commands. To change the container I use the `.ko.yaml` and include the following.

```yaml
defaultBaseImage: gcr.io/distroless/static:debug
```

* Enabling `--bare` to simplify the path the container is published under, the default naming structure isn't supported by ECR.

To run `ko` I use the command as such:

```
ko publish --platform=linux/arm64 --image-label arch=arm64 --image-label git_hash=abc123 --bare ./cmd/api-lambda
```

Take a look in the `Makefile` for more details of how I configure these parameters.

# Deploying

## Prerequisites

1. You will need to export an `AWS_PROFILE` and `AWS_DEFAULT_REGION` to enable access to AWS.
2. You will need to export `SAM_BUCKET` environment variable which contains the name of an S3 bucket in the same region as your Deploying.

To manage these environment variables you can create an `.envrc` using the `.envrc.example` and update it with your settings, this is used with [direnv](https://direnv.net/).

```
cp .envrc.example .envrc
```

Then modify these vars in this example file.

# Commands

To deploy the ECR repository template.

```
make deploy-repository
```

To login to docker using the ECR repository.

```
make docker-login
```

To deploy the API.

```
make deploy-api
```

# License

This application is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au).