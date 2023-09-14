"# Golang-psql-k8s-Banking-app" 
Tools 
DB diagram.io
Table plus as a GUI for postgres
# setting The DB 
docker pull postgres:12-alpine
docker run --name postgres12 -e POSTGRES_PASSWORD=postgres  -e POSTGRES_USER=root -d -p 5432:5432 postgres:12-alpine


# Running Migrations 
migrate -path db/migration -database "" -verbose up 