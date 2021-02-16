package main

type customerService interface {
	getCustomer(id string) Customer
	getOrders(customerId string) []Order
}
type customerServiceImpl struct {
 CDao customerDao `inject:""`
 ODao orderDao `inject:""`
}
func (c customerServiceImpl) getCustomer(id string) Customer {
	return c.CDao.getCustomer(id)
}
func (c customerServiceImpl) getOrders(customerId string) []Order {
	return c.ODao.getOrders(customerId)
}
type Customer struct {
	Id string `json:"id"`
	Name string `json:"name"`
}