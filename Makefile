
create-network:
	docker network create -d bridge mongorm

up:
	docker-compose -f tools/docker-compose.test.yml up #-d

down:
	docker-compose -f tools/docker-compose.test.yml down