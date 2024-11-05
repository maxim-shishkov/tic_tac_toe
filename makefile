IMAGE_NAME := tic_tac_toe
TAG := latest
CONTAINER_NAME := tic_tac_toe

.PHONY: up down qwe

up:
	docker build -t $(IMAGE_NAME):$(TAG) -f ./docker/app/dockerfile .
	docker run -d --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME):$(TAG)

down:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)
