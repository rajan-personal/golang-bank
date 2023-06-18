Docker Commands

1. docker pull postgres:16beta1-alpine
2. docker run --name postgresql-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=somePassword -d postgres:16beta1-alpine
3. docker exec -it postgresql-db psql -U root
