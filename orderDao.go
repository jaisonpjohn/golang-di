package main

type orderDao interface {
	getOrders(customerId string) []Order
}
type orderDaoImpl struct {

}
type Order struct {
id string
name string
customerId string
}

func (orderDaoImpl) getOrders(customerId string) []Order  {
	return []Order{{id: "123",name: "some Order1", customerId: customerId}}
}