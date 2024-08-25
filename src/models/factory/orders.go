// Code generated by BobGen psql v0.28.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"context"
	"testing"
	"time"

	"github.com/aarondl/opt/null"
	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/jaswdr/faker/v2"
	"github.com/stephenafamo/bob"
	models "pickle_store/models"
)

type OrderMod interface {
	Apply(*OrderTemplate)
}

type OrderModFunc func(*OrderTemplate)

func (f OrderModFunc) Apply(n *OrderTemplate) {
	f(n)
}

type OrderModSlice []OrderMod

func (mods OrderModSlice) Apply(n *OrderTemplate) {
	for _, f := range mods {
		f.Apply(n)
	}
}

// OrderTemplate is an object representing the database table.
// all columns are optional and should be set by mods
type OrderTemplate struct {
	OrderID    func() int32
	CustomerID func() null.Val[int32]
	ProductID  func() null.Val[int32]
	Quantity   func() int32
	OrderDate  func() null.Val[time.Time]

	r orderR
	f *Factory
}

type orderR struct {
	Customer *orderRCustomerR
	Product  *orderRProductR
}

type orderRCustomerR struct {
	o *CustomerTemplate
}
type orderRProductR struct {
	o *ProductTemplate
}

// Apply mods to the OrderTemplate
func (o *OrderTemplate) Apply(mods ...OrderMod) {
	for _, mod := range mods {
		mod.Apply(o)
	}
}

// toModel returns an *models.Order
// this does nothing with the relationship templates
func (o OrderTemplate) toModel() *models.Order {
	m := &models.Order{}

	if o.OrderID != nil {
		m.OrderID = o.OrderID()
	}
	if o.CustomerID != nil {
		m.CustomerID = o.CustomerID()
	}
	if o.ProductID != nil {
		m.ProductID = o.ProductID()
	}
	if o.Quantity != nil {
		m.Quantity = o.Quantity()
	}
	if o.OrderDate != nil {
		m.OrderDate = o.OrderDate()
	}

	return m
}

// toModels returns an models.OrderSlice
// this does nothing with the relationship templates
func (o OrderTemplate) toModels(number int) models.OrderSlice {
	m := make(models.OrderSlice, number)

	for i := range m {
		m[i] = o.toModel()
	}

	return m
}

// setModelRels creates and sets the relationships on *models.Order
// according to the relationships in the template. Nothing is inserted into the db
func (t OrderTemplate) setModelRels(o *models.Order) {
	if t.r.Customer != nil {
		rel := t.r.Customer.o.toModel()
		rel.R.Orders = append(rel.R.Orders, o)
		o.CustomerID = null.From(rel.CustomerID)
		o.R.Customer = rel
	}

	if t.r.Product != nil {
		rel := t.r.Product.o.toModel()
		rel.R.Orders = append(rel.R.Orders, o)
		o.ProductID = null.From(rel.ProductID)
		o.R.Product = rel
	}
}

// BuildSetter returns an *models.OrderSetter
// this does nothing with the relationship templates
func (o OrderTemplate) BuildSetter() *models.OrderSetter {
	m := &models.OrderSetter{}

	if o.OrderID != nil {
		m.OrderID = omit.From(o.OrderID())
	}
	if o.CustomerID != nil {
		m.CustomerID = omitnull.FromNull(o.CustomerID())
	}
	if o.ProductID != nil {
		m.ProductID = omitnull.FromNull(o.ProductID())
	}
	if o.Quantity != nil {
		m.Quantity = omit.From(o.Quantity())
	}
	if o.OrderDate != nil {
		m.OrderDate = omitnull.FromNull(o.OrderDate())
	}

	return m
}

// BuildManySetter returns an []*models.OrderSetter
// this does nothing with the relationship templates
func (o OrderTemplate) BuildManySetter(number int) []*models.OrderSetter {
	m := make([]*models.OrderSetter, number)

	for i := range m {
		m[i] = o.BuildSetter()
	}

	return m
}

// Build returns an *models.Order
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use OrderTemplate.Create
func (o OrderTemplate) Build() *models.Order {
	m := o.toModel()
	o.setModelRels(m)

	return m
}

// BuildMany returns an models.OrderSlice
// Related objects are also created and placed in the .R field
// NOTE: Objects are not inserted into the database. Use OrderTemplate.CreateMany
func (o OrderTemplate) BuildMany(number int) models.OrderSlice {
	m := make(models.OrderSlice, number)

	for i := range m {
		m[i] = o.Build()
	}

	return m
}

func ensureCreatableOrder(m *models.OrderSetter) {
	if m.Quantity.IsUnset() {
		m.Quantity = omit.From(random_int32(nil))
	}
}

// insertOptRels creates and inserts any optional the relationships on *models.Order
// according to the relationships in the template.
// any required relationship should have already exist on the model
func (o *OrderTemplate) insertOptRels(ctx context.Context, exec bob.Executor, m *models.Order) (context.Context, error) {
	var err error

	if o.r.Customer != nil {
		var rel0 *models.Customer
		ctx, rel0, err = o.r.Customer.o.create(ctx, exec)
		if err != nil {
			return ctx, err
		}
		err = m.AttachCustomer(ctx, exec, rel0)
		if err != nil {
			return ctx, err
		}
	}

	if o.r.Product != nil {
		var rel1 *models.Product
		ctx, rel1, err = o.r.Product.o.create(ctx, exec)
		if err != nil {
			return ctx, err
		}
		err = m.AttachProduct(ctx, exec, rel1)
		if err != nil {
			return ctx, err
		}
	}

	return ctx, err
}

// Create builds a order and inserts it into the database
// Relations objects are also inserted and placed in the .R field
func (o *OrderTemplate) Create(ctx context.Context, exec bob.Executor) (*models.Order, error) {
	_, m, err := o.create(ctx, exec)
	return m, err
}

// MustCreate builds a order and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// panics if an error occurs
func (o *OrderTemplate) MustCreate(ctx context.Context, exec bob.Executor) *models.Order {
	_, m, err := o.create(ctx, exec)
	if err != nil {
		panic(err)
	}
	return m
}

// CreateOrFail builds a order and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// It calls `tb.Fatal(err)` on the test/benchmark if an error occurs
func (o *OrderTemplate) CreateOrFail(ctx context.Context, tb testing.TB, exec bob.Executor) *models.Order {
	tb.Helper()
	_, m, err := o.create(ctx, exec)
	if err != nil {
		tb.Fatal(err)
		return nil
	}
	return m
}

// create builds a order and inserts it into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted model
func (o *OrderTemplate) create(ctx context.Context, exec bob.Executor) (context.Context, *models.Order, error) {
	var err error
	opt := o.BuildSetter()
	ensureCreatableOrder(opt)

	m, err := models.Orders.Insert(ctx, exec, opt)
	if err != nil {
		return ctx, nil, err
	}
	ctx = orderCtx.WithValue(ctx, m)

	ctx, err = o.insertOptRels(ctx, exec, m)
	return ctx, m, err
}

// CreateMany builds multiple orders and inserts them into the database
// Relations objects are also inserted and placed in the .R field
func (o OrderTemplate) CreateMany(ctx context.Context, exec bob.Executor, number int) (models.OrderSlice, error) {
	_, m, err := o.createMany(ctx, exec, number)
	return m, err
}

// MustCreateMany builds multiple orders and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// panics if an error occurs
func (o OrderTemplate) MustCreateMany(ctx context.Context, exec bob.Executor, number int) models.OrderSlice {
	_, m, err := o.createMany(ctx, exec, number)
	if err != nil {
		panic(err)
	}
	return m
}

// CreateManyOrFail builds multiple orders and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// It calls `tb.Fatal(err)` on the test/benchmark if an error occurs
func (o OrderTemplate) CreateManyOrFail(ctx context.Context, tb testing.TB, exec bob.Executor, number int) models.OrderSlice {
	tb.Helper()
	_, m, err := o.createMany(ctx, exec, number)
	if err != nil {
		tb.Fatal(err)
		return nil
	}
	return m
}

// createMany builds multiple orders and inserts them into the database
// Relations objects are also inserted and placed in the .R field
// this returns a context that includes the newly inserted models
func (o OrderTemplate) createMany(ctx context.Context, exec bob.Executor, number int) (context.Context, models.OrderSlice, error) {
	var err error
	m := make(models.OrderSlice, number)

	for i := range m {
		ctx, m[i], err = o.create(ctx, exec)
		if err != nil {
			return ctx, nil, err
		}
	}

	return ctx, m, nil
}

// Order has methods that act as mods for the OrderTemplate
var OrderMods orderMods

type orderMods struct{}

func (m orderMods) RandomizeAllColumns(f *faker.Faker) OrderMod {
	return OrderModSlice{
		OrderMods.RandomOrderID(f),
		OrderMods.RandomCustomerID(f),
		OrderMods.RandomProductID(f),
		OrderMods.RandomQuantity(f),
		OrderMods.RandomOrderDate(f),
	}
}

// Set the model columns to this value
func (m orderMods) OrderID(val int32) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.OrderID = func() int32 { return val }
	})
}

// Set the Column from the function
func (m orderMods) OrderIDFunc(f func() int32) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.OrderID = f
	})
}

// Clear any values for the column
func (m orderMods) UnsetOrderID() OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.OrderID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m orderMods) RandomOrderID(f *faker.Faker) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.OrderID = func() int32 {
			return random_int32(f)
		}
	})
}

// Set the model columns to this value
func (m orderMods) CustomerID(val null.Val[int32]) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.CustomerID = func() null.Val[int32] { return val }
	})
}

// Set the Column from the function
func (m orderMods) CustomerIDFunc(f func() null.Val[int32]) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.CustomerID = f
	})
}

// Clear any values for the column
func (m orderMods) UnsetCustomerID() OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.CustomerID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m orderMods) RandomCustomerID(f *faker.Faker) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.CustomerID = func() null.Val[int32] {
			if f == nil {
				f = &defaultFaker
			}

			if f.Bool() {
				return null.FromPtr[int32](nil)
			}

			return null.From(random_int32(f))
		}
	})
}

// Set the model columns to this value
func (m orderMods) ProductID(val null.Val[int32]) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.ProductID = func() null.Val[int32] { return val }
	})
}

// Set the Column from the function
func (m orderMods) ProductIDFunc(f func() null.Val[int32]) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.ProductID = f
	})
}

// Clear any values for the column
func (m orderMods) UnsetProductID() OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.ProductID = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m orderMods) RandomProductID(f *faker.Faker) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.ProductID = func() null.Val[int32] {
			if f == nil {
				f = &defaultFaker
			}

			if f.Bool() {
				return null.FromPtr[int32](nil)
			}

			return null.From(random_int32(f))
		}
	})
}

// Set the model columns to this value
func (m orderMods) Quantity(val int32) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.Quantity = func() int32 { return val }
	})
}

// Set the Column from the function
func (m orderMods) QuantityFunc(f func() int32) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.Quantity = f
	})
}

// Clear any values for the column
func (m orderMods) UnsetQuantity() OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.Quantity = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m orderMods) RandomQuantity(f *faker.Faker) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.Quantity = func() int32 {
			return random_int32(f)
		}
	})
}

// Set the model columns to this value
func (m orderMods) OrderDate(val null.Val[time.Time]) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.OrderDate = func() null.Val[time.Time] { return val }
	})
}

// Set the Column from the function
func (m orderMods) OrderDateFunc(f func() null.Val[time.Time]) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.OrderDate = f
	})
}

// Clear any values for the column
func (m orderMods) UnsetOrderDate() OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.OrderDate = nil
	})
}

// Generates a random value for the column using the given faker
// if faker is nil, a default faker is used
func (m orderMods) RandomOrderDate(f *faker.Faker) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.OrderDate = func() null.Val[time.Time] {
			if f == nil {
				f = &defaultFaker
			}

			if f.Bool() {
				return null.FromPtr[time.Time](nil)
			}

			return null.From(random_time_Time(f))
		}
	})
}

func (m orderMods) WithCustomer(rel *CustomerTemplate) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.r.Customer = &orderRCustomerR{
			o: rel,
		}
	})
}

func (m orderMods) WithNewCustomer(mods ...CustomerMod) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		related := o.f.NewCustomer(mods...)

		m.WithCustomer(related).Apply(o)
	})
}

func (m orderMods) WithoutCustomer() OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.r.Customer = nil
	})
}

func (m orderMods) WithProduct(rel *ProductTemplate) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.r.Product = &orderRProductR{
			o: rel,
		}
	})
}

func (m orderMods) WithNewProduct(mods ...ProductMod) OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		related := o.f.NewProduct(mods...)

		m.WithProduct(related).Apply(o)
	})
}

func (m orderMods) WithoutProduct() OrderMod {
	return OrderModFunc(func(o *OrderTemplate) {
		o.r.Product = nil
	})
}