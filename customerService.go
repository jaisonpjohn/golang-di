package main

type customerService interface {
	getCustomer(id string) Customer
	getOrders(customerId string) Order
}
type customerServiceImpl struct {
 cDao *customerDao `inject`
 oDao *orderDao `inject`
}
func (c customerServiceImpl) getCustomer(id string) Customer {
	dao := *c.cDao
	return dao.getCustomer(id)
}
func (c customerServiceImpl) getOrders(customerId string) []Order {
	dao := *c.oDao
	return dao.getOrders(customerId)
}
type Customer struct {
	id string
	name string
}