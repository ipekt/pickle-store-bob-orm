// Code generated by BobGen psql v0.28.1. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

type Factory struct {
	baseCustomerMods CustomerModSlice
	baseOrderMods    OrderModSlice
	baseProductMods  ProductModSlice
}

func New() *Factory {
	return &Factory{}
}

func (f *Factory) NewCustomer(mods ...CustomerMod) *CustomerTemplate {
	o := &CustomerTemplate{f: f}

	if f != nil {
		f.baseCustomerMods.Apply(o)
	}

	CustomerModSlice(mods).Apply(o)

	return o
}

func (f *Factory) NewOrder(mods ...OrderMod) *OrderTemplate {
	o := &OrderTemplate{f: f}

	if f != nil {
		f.baseOrderMods.Apply(o)
	}

	OrderModSlice(mods).Apply(o)

	return o
}

func (f *Factory) NewProduct(mods ...ProductMod) *ProductTemplate {
	o := &ProductTemplate{f: f}

	if f != nil {
		f.baseProductMods.Apply(o)
	}

	ProductModSlice(mods).Apply(o)

	return o
}

func (f *Factory) ClearBaseCustomerMods() {
	f.baseCustomerMods = nil
}

func (f *Factory) AddBaseCustomerMod(mods ...CustomerMod) {
	f.baseCustomerMods = append(f.baseCustomerMods, mods...)
}

func (f *Factory) ClearBaseOrderMods() {
	f.baseOrderMods = nil
}

func (f *Factory) AddBaseOrderMod(mods ...OrderMod) {
	f.baseOrderMods = append(f.baseOrderMods, mods...)
}

func (f *Factory) ClearBaseProductMods() {
	f.baseProductMods = nil
}

func (f *Factory) AddBaseProductMod(mods ...ProductMod) {
	f.baseProductMods = append(f.baseProductMods, mods...)
}