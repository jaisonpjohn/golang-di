package main

type customerService interface {
	getCustomer(id string) Customer
	getOrders(customerId string) []Order
}
type customerServiceImpl struct {
 cDao customerDao
 oDao orderDao
}
func (c customerServiceImpl) getCustomer(id string) Customer {
	return c.cDao.getCustomer(id)
}
func (c customerServiceImpl) getOrders(customerId string) []Order {
	return c.oDao.getOrders(customerId)
}
type Customer struct {
	Id string `json:"id"`
	Name string `json:"name"`
}