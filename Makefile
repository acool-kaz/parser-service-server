run:
	go run cmd/main.go

open:
	docker exec -it service_db psql -U postgres

start:
	docker run --name=service_db -e POSTGRES_PASSWORD='12345' -p 5436:5436 -d --rm postgres
	@sleep 5
	go run cmd/main.go

restart:
	docker stop service_db
	docker run --name=service_db -e POSTGRES_PASSWORD='12345' -p 5436:5436 -d --rm postgres
	@sleep 5
	go run cmd/main.go

gen:
	protoc --go_out=. \
		--go-grpc_out=. \
		./proto/parser.proto