// Code generated by BobGen psql v0.28.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"
	"testing"

	"github.com/aarondl/opt/null"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/jaswdr/faker/v2"
	"github.com/shopspring/decimal"
	"github.com/stephenafamo/bob"
	models "pickle_store/models"
)

type ProductMod interface {
	Apply(*ProductTemplate)
}

type ProductModFunc func(*ProductTemplate)

func (f ProductModFunc) Apply(n *ProductTemplate) {
	f(n)
}

type ProductModSlice []ProductMod

func (mods ProductModSlice) Apply(n *ProductTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// ProductTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type ProductTemplate struct {
	ProductID     func() int32
	Name          func() string
	Description   func() null.Val[string]
	Price         func() decimal.Decimal
	StockQuantity func() int32

	r productR
	f *Factory
}

type productR struct {
	Orders []*productROrdersR
}

type productROrdersR struct {
	number int
	o      *OrderTemplate
}

// Apply mods to the ProductTemplate
func (o *ProductTemplate) Apply(mods ...ProductMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.Product
// this does nothing with the relationship templates
func (o ProductTemplate) toModel() *models.Product {
	m := &models.Product{}

	if o.ProductID != nil {
		m.ProductID = o.ProductID()
	}
	if o.Name != nil {
		m.Name = o.Name()
	}
	if o.Description != nil {
		m.Description = o.Description()
	}
	if o.Price != nil {
		m.Price = o.Price()
	}
	if o.StockQuantity != nil {
		m.StockQuantity = o.StockQuantity()
	}

	return m
}

// toModels returns an models.ProductSlice
// this does nothing with the relationship templates
func (o ProductTemplate) toModels(number int) models.ProductSlice {
	m := make(models.ProductSlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.Product
// according to the relationships in the template. Nothing is inserted into the db
func (t ProductTemplate) setModelRels(o *models.Product) {
	if t.r.Orders != nil {
		rel := models.OrderSlice{}
		for _, r := range t.r.Orders {
			related := r.o.toModels(r.number)
			for _, rel := range related {
				rel.ProductID = null.From(o.ProductID)
				rel.R.Product = o
			}
			rel = append(rel, related...)
		}
		o.R.Orders = rel
	}
}

// BuildSetter returns an *models.ProductSetter
// this does nothing with the relationship templates
func (o ProductTemplate) BuildSetter() *models.ProductSetter {
	m := &models.ProductSetter{}

	if o.ProductID != nil {
		m.ProductID = omit.From(o.ProductID())
	}
	if o.Name != nil {
		m.Name = omit.From(o.Name())
	}
	if o.Description != nil {
		m.Description = omitnull.FromNull(o.Description())
	}
	if o.Price != nil {
		m.Price = omit.From(o.Price())
	}
	if o.StockQuantity != nil {
		m.StockQuantity = omit.From(o.StockQuantity())
	}

	return m
}

// BuildManySetter returns an []*models.ProductSetter
// this does nothing with the relationship templates
func (o ProductTemplate) BuildManySetter(number int) []*models.ProductSetter {
	m := make([]*models.ProductSetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.Product
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use ProductTemplate.Create
func (o ProductTemplate) Build() *models.Product {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.ProductSlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use ProductTemplate.CreateMany
func (o ProductTemplate) BuildMany(number int) models.ProductSlice {
	m := make(models.ProductSlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatableProduct(m *models.ProductSetter) {
	if m.Name.IsUnset() {
		m.Name = omit.From(random_string(nil))
	}
	if m.Price.IsUnset() {
		m.Price = omit.From(random_decimal_Decimal(nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.Product
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *ProductTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.Product) (context.Context, error) {
	var err error

	if o.r.Orders != nil {
		for _, r := range o.r.Orders {
			var rel0 models.OrderSlice
			ctx, rel0, err = r.o.createMany(ctx, exec, r.number)
			if err != nil {
				return ctx, err
			}

			err = m.AttachOrders(ctx, exec, rel0...)
			if err != nil {
				return ctx, err
			}
		}
	}

	return ctx, err
}

// Create builds a product and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *ProductTemplate) Create(ctx context.Context, exec bob.Executor) (*models.Product, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// MustCreate builds a product and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// panics if an error occurs
func (o *ProductTemplate) MustCreate(ctx context.Context, exec bob.Executor) *models.Product {
	_, m, err := o.create(ctx, exec)
	if err != nil {
		panic(err)
	}
	return m
}

// CreateOrFail builds a product and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// It calls `tb.Fatal(err)` on the test/benchmark if an error occurs
func (o *ProductTemplate) CreateOrFail(ctx context.Context, tb testing.TB, exec bob.Executor) *models.Product {
	tb.Helper()
	_, m, err := o.create(ctx, exec)
	if err != nil {
		tb.Fatal(err)
		return nil
	}
	return m
}

// create builds a product and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *ProductTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.Product, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatableProduct(opt)

	m, err := models.Products.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = productCtx.WithValue(ctx, m)

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple products and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o ProductTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.ProductSlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// MustCreateMany builds multiple products and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// panics if an error occurs
func (o ProductTemplate) MustCreateMany(ctx context.Context, exec bob.Executor, number int) models.ProductSlice {
	_, m, err := o.createMany(ctx, exec, number)
	if err != nil {
		panic(err)
	}
	return m
}

// CreateManyOrFail builds multiple products and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// It calls `tb.Fatal(err)` on the test/benchmark if an error occurs
func (o ProductTemplate) CreateManyOrFail(ctx context.Context, tb testing.TB, exec bob.Executor, number int) models.ProductSlice {
	tb.Helper()
	_, m, err := o.createMany(ctx, exec, number)
	if err != nil {
		tb.Fatal(err)
		return nil
	}
	return m
}

// createMany builds multiple products and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o ProductTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.ProductSlice, error) {
	var err error
	m := make(models.ProductSlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// Product has methods that act as mods for the ProductTemplate
var ProductMods productMods

type productMods struct{}

func (m productMods) RandomizeAllColumns(f *faker.Faker) ProductMod {
	return ProductModSlice{
		ProductMods.RandomProductID(f),
		ProductMods.RandomName(f),
		ProductMods.RandomDescription(f),
		ProductMods.RandomPrice(f),
		ProductMods.RandomStockQuantity(f),
	}
}

// Set the model columns to this value
func (m productMods) ProductID(val int32) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.ProductID = func() int32 { return val }
	})
}

// Set the Column from the function
func (m productMods) ProductIDFunc(f func() int32) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.ProductID = f
	})
}

// Clear any values for the column
func (m productMods) UnsetProductID() ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.ProductID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m productMods) RandomProductID(f *faker.Faker) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.ProductID = func() int32 {
			return random_int32(f)
		}
	})
}

// Set the model columns to this value
func (m productMods) Name(val string) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Name = func() string { return val }
	})
}

// Set the Column from the function
func (m productMods) NameFunc(f func() string) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Name = f
	})
}

// Clear any values for the column
func (m productMods) UnsetName() ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Name = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m productMods) RandomName(f *faker.Faker) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Name = func() string {
			return random_string(f)
		}
	})
}

// Set the model columns to this value
func (m productMods) Description(val null.Val[string]) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Description = func() null.Val[string] { return val }
	})
}

// Set the Column from the function
func (m productMods) DescriptionFunc(f func() null.Val[string]) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Description = f
	})
}

// Clear any values for the column
func (m productMods) UnsetDescription() ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Description = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m productMods) RandomDescription(f *faker.Faker) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Description = func() null.Val[string] {
			if f == nil {
				f = &defaultFaker
			}

			if f.Bool() {
				return null.FromPtr[string](nil)
			}

			return null.From(random_string(f))
		}
	})
}

// Set the model columns to this value
func (m productMods) Price(val decimal.Decimal) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Price = func() decimal.Decimal { return val }
	})
}

// Set the Column from the function
func (m productMods) PriceFunc(f func() decimal.Decimal) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Price = f
	})
}

// Clear any values for the column
func (m productMods) UnsetPrice() ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Price = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m productMods) RandomPrice(f *faker.Faker) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.Price = func() decimal.Decimal {
			return random_decimal_Decimal(f)
		}
	})
}

// Set the model columns to this value
func (m productMods) StockQuantity(val int32) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.StockQuantity = func() int32 { return val }
	})
}

// Set the Column from the function
func (m productMods) StockQuantityFunc(f func() int32) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.StockQuantity = f
	})
}

// Clear any values for the column
func (m productMods) UnsetStockQuantity() ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.StockQuantity = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m productMods) RandomStockQuantity(f *faker.Faker) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.StockQuantity = func() int32 {
			return random_int32(f)
		}
	})
}

func (m productMods) WithOrders(number int, related *OrderTemplate) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.r.Orders = []*productROrdersR{{
			number: number,
			o:      related,
		}}
	})
}

func (m productMods) WithNewOrders(number int, mods ...OrderMod) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		related := o.f.NewOrder(mods...)
		m.WithOrders(number, related).Apply(o)
	})
}

func (m productMods) AddOrders(number int, related *OrderTemplate) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.r.Orders = append(o.r.Orders, &productROrdersR{
			number: number,
			o:      related,
		})
	})
}

func (m productMods) AddNewOrders(number int, mods ...OrderMod) ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		related := o.f.NewOrder(mods...)
		m.AddOrders(number, related).Apply(o)
	})
}

func (m productMods) WithoutOrders() ProductMod {
	return ProductModFunc(func(o *ProductTemplate) {
		o.r.Orders = nil
	})
}
