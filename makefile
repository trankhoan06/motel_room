DB_URL=mysql://admin:Khanh123456@b3zg7i71qhr2grdmjjsz-mysql.services.clever-cloud.com:3306/b3zg7i71qhr2grdmjjsz

network:
	docker network create bank-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

mysql:
	docker run --name project_personal -e MYSQL_ROOT_PASSWORD=123456 -d -p 3309:3306 mysql:8.0.31

createdb:
	docker exec -it postgres createdb --username=root --owner=root project_personal

dropdb:
	docker exec -it postgres dropdb simple_bank

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine
test:
	go test -v -cover -short ./...

start:
	go run main.go

reown:
	sudo chown -R $(u) .

.PHONY: run migrate_up migrate_down new_migration test start createdb dropdb postgres mysql network reown redis
