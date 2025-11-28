APP_NAME=telegram-bot
REGISTRY=ghcr.io
REPOSITORY=samtishka/telegram-bot
TAG=v1.0.0-linux-amd64
IMAGE=$(REGISTRY)/$(REPOSITORY):$(TAG)
PLATFORM=linux/amd64

.PHONY: build docker-build docker-push helm-package

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bot main.go

docker-build:
	docker build --platform $(PLATFORM) -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)

helm-package:
	helm lint telegram-bot
	helm package telegram-bot
