migrateup:
migrate -path migrations -database "postgres://postgres:lbfc2005@localhost:5432/d.ibragimovDB?sslmode=disable" -verbose up
migratedown:
migrate -path migrations -database "postgres://postgres:lbfc2005@localhost:5432/d.ibragimovDB?sslmode=disable" -verbose down