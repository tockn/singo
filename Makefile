.PHONY: sdk-build docker-build docker-push

sdk-build:
	cd sdk;yarn build

docker-build:
	docker build -t tockn/singo .

docker-push:
	docker push tockn/singo

all: sdk-build docker-build docker-push

up:
	docker-compose up --build

sdk-dev:
	cd sdk; yarn start
