.PHONY: postgres migrate run tidy

build-banking:
	docker build -t banking-service ./banking-service

build-credit:
	docker build -t credit-search ./credit-search-service

build-community:
	docker build -t community-service ./community-service

build-all: build-banking build-credit build-community

up:
	docker-compose up

up-detatched:
	docker-compose up -d