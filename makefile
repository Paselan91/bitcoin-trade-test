.PHONY: init plan apply destroy check
ps:
	docker-compose ps
attach app:
	docker-compose exec bitflyer bash
attach db:
	docker-compose exec mysql bash
fmt:
	docker exec -it be go fmt ./...
lint:
	docker exec -it fe yarn lint --fix .
format:
	docker exec -it fe yarn lint --fix .
	docker exec -it be go fmt ./...
up:
	docker-compose up
up-d:
	docker-compose up -d
restart:
	docker-compose restart
down:
	docker-compose down
