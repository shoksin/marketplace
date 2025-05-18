start:
	docker-compose up
start-rebuild:
	docker-compose up --build
stop:
	docker-compose down
open-auth-db:
	docker exec -it marketplace_auth-db_1 psql -U postgres -d auth-service
clean-images:
	docker rmi $(shell docker images -q)
clean-containers:
	docker rm $(shell docker ps -aq)