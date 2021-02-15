package main

type customerDao interface {
	getCustomer(id string) Customer
}
type customerDaoImpl struct {

}

func (customerDaoImpl) getCustomer(id string) Customer {
	return Customer{Id: id, Name: "John Doe"}
}