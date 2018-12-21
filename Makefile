PROJECT_DIRECTORY ?= "pulumi-poc"
STACK ?= "pulumi-poc-dev"
CONTAINER_GCP_NAME = "pulumi:gcp"

#Init is for doc propouse, it is only need once when you are creating the project
init:
	docker run -ti -e PULUMI_ACCESS_TOKEN="$$PULUMI_ACCESS_TOKEN" -v $$PWD:/app -ti pulumi/actions new gcp-go --dir $(PROJECT_DIRECTORY)
build-gcp:
	docker build . -t $(CONTAINER_GCP_NAME)
pulumi-update:
	docker run -ti -e PULUMI_ACCESS_TOKEN="$$PULUMI_ACCESS_TOKEN" -e GREETING="$$GREETING" -e GOOGLE_PROJECT="$$GOOGLE_PROJECT" -e GOOGLE_CREDENTIALS="$$CREDENTIALS" $(CONTAINER_GCP_NAME) update -s $(STACK)
pulumi-preview:
	docker run -ti -e PULUMI_ACCESS_TOKEN="$$PULUMI_ACCESS_TOKEN" -e GREETING="$$GREETING" -e GOOGLE_PROJECT="$$GOOGLE_PROJECT" -e GOOGLE_CREDENTIALS="$$CREDENTIALS"  $(CONTAINER_GCP_NAME) preview -s $(STACK)
pulumi-destroy:
	docker run -ti -e PULUMI_ACCESS_TOKEN="$$PULUMI_ACCESS_TOKEN" -e GREETING="$$GREETING" -e GOOGLE_PROJECT="$$GOOGLE_PROJECT" -e GOOGLE_CREDENTIALS="$$CREDENTIALS" $(CONTAINER_GCP_NAME) destroy -s $(STACK)
#Due a bug using ENV vars described in https://github.com/pulumi/pulumi-gcp/issues/25 for 0.15.1
config:
	docker run -ti -e PULUMI_ACCESS_TOKEN="$$PULUMI_ACCESS_TOKEN"-v $$PWD:/app  -e GOOGLE_CREDENTIALS="$$CREDENTIALS" pulumi:gcp config set gcp:project
