postgres:
	docker run --name postgres -e POSTGRES_UsER=postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:12-alpine

createdb:
	createdb --username=postgres --owner=postgres  simple_bank
	@echo "Database created"
   

dropdb:
	dropdb  --username=postgres simple_bank	
	@echo "Database dropped"

migrateup:
	migrate -path db/migration -database "postgresql://postgres:5463@localhost:5432/simple_bank?sslmode=disable" -verbose   force 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:5463@localhost:5432/simple_bank?sslmode=disable" -verbose force 1
	
.PHONY: createdb dropdb migrateup migratedown


