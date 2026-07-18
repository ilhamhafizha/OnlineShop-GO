# Online Shop Project 

1. Jalankan Docker untuk PostgreSQL

```
docker run --name postgresql -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=database -d -p 5432:5432 postgres:16
``` 

2. Export Enviroment yang dibutuhkan

```
export DB_URI=postgresql://user:password@localhost:5432/database?sslmode=disable
export ADMIN_SECRET=secret
```

3. Jalankan Program

```
go run main.go
```