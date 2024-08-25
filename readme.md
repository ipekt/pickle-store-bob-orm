# Pickle Store

Example code using [bob orm](https://github.com/stephenafamo/bob)

## Setup db

```sh
psql -U admin -h localhost -p 5432 -d pickles_db

-- copy paste init.sql
```

## How to generate bob code

```sh
PSQL_DSN=postgres://admin:root@localhost:5432/pickles_db?sslmode=disable go run github.com/stephenafamo/bob/gen/bobgen-psql@latest
```


## Run

```sh
go run ./src/main.go
```

## Query examples

### Left join

```go
	q := psql.Select(
		sm.Columns(
			psql.Quote("products", "product_id"),
			psql.Quote("products", "name"),
			psql.Quote("o", "customer_id")),
		sm.From("products"),
		sm.LeftJoin(psql.Quote("orders").As("o")).OnEQ(psql.Quote("products", "product_id"), psql.Quote("o", "product_id")),
	)

```
