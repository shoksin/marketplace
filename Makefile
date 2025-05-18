DIRS = api-gateway auth order product

start:
	docker-compose up
start-rebuild:
	docker-compose up --build
stop:
	docker-compose down
open-auth-db:
	docker exec -it marketplace_auth-db_1 psql -U postgres -d auth-service
open-product-db:
	docker exec -it marketplace_product-db_1 psql -U postgres product-service
clean-images:
	docker rmi $(shell docker images -q)
clean-containers:
	docker rm $(shell docker ps -aq)
get-last-proto:
	for dir in $(DIRS); do \
    		echo "Выполнение в директории: $$dir"; \
    		( cd $$dir && go get github.com/shoksin/marketplace-protos@v0.0.9 ); \
    	done