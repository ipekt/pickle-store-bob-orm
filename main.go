package main

import (
	"context"
	"fmt"

	"github.com/ipekt/pickle-store/src/models"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/stephenafamo/bob"

	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/scan"
)

type Product struct {
	ProductID     string
	Name          string
	CustomerID    string
	TotalQuantity int
}

func main() {
	// PostgreSQL connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "admin", "root", "pickles_db")

	db, err := bob.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	debugdb := bob.Debug(db)

	defer db.Close()

	// Lateral join example
	q := psql.Select(
		sm.Columns(
			models.ColumnNames.Products.ProductID,
			models.ColumnNames.Products.Name,
			psql.Quote("count", "total_quantity"),
		),
		sm.From(models.TableNames.Products),
		sm.LeftJoin(
			psql.RawQuery("SELECT SUM(quantity) AS total_quantity FROM orders WHERE product_id = products.product_id"),
		).
			Lateral().
			As("count").
			On(psql.Arg(true)),
	)

	p, err := bob.All(context.Background(), debugdb, q, scan.StructMapper[Product]())
	if err != nil {
		logrus.Errorf("exec fail %+v", err)
	}
	fmt.Print(p)
}
