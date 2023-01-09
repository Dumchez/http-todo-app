# REST API

## It allows users to sign up and login, also to create/update/delete lists & items

### To Run The App:

```
docker run --name=todo-db -e POSTGRES_PASSWORD='your_password' -p 5436:5432 -d --rm postgres
migrate -path ./schema -database 'postgres://postgres:your_password@localhost:5436/postgres?sslmode=disable' up
```
