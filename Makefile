postgres:
	docker run --name postgres -e POSTGRES_UsER=postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:12-alpine

createdb:
	createdb --username=postgres --owner=postgres  simple_bank
	@echo "Database created"
   

dropdb:
	dropdb  --username=postgres simple_bank	
	@echo "Database dropped"

migrateup:
	migrate -path db/migration -database "postgresql://postgres:5463@localhost:5432/simple_bank?sslmode=disable" -verbose up    


migratedown:
	migrate -path db/migration -database "postgresql://postgres:5463@localhost:5432/simple_bank?sslmode=disable" -verbose  down
	
sqlc:
	docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate


.PHONY: createdb dropdb migrateup migratedown sqlc



