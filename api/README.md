# SPAJZA backend

We have several scripts to help:

## Load google category

```sh
go run ./cmd/category_seeder/main.go -pwd $DB_PWD
```

## Reload schema from DB to sqlc.sql

```sh
go run ./cmd/schema_loader/main.go -pwd $DB_PWD
```

## Generate DB loaders

First we need to instal sqlc dependency by `brew install sqlc`

```sh
sqlc
```

## Jet DB

```sh
jet -source=mysql -host=localhost -port=3306 -user=user -password=user -dbname=zasobar -path=./adapters/repository
```
