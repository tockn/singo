.PHONY: sdk-build docker-clean docker-run docker-run-example docker-build docker-push all up sdk-dev

docker-clean:
	-docker rm singo

docker-run: docker-clean
	docker run -p 5000:5000 --name singo tockn/singo:latest

docker-run-example: docker-clean
	docker run -p 5000:5000 --name singo tockn/singo:latest --example

docker-build:
	docker build -t tockn/singo .

docker-push:
	docker push tockn/singo

all: sdk-build docker-build docker-push

up:
	docker-compose up --build

build:
	go build -o singo

sdk-build:
	cd sdk;yarn build

sdk-dev:
	cd sdk; yarn start

run: build
	./singo

run-example: build sdk-build
	./singo -example
