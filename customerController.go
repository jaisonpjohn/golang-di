package main


type customerController struct {
 customerService *customerService `inject`
}

func (c customerController) getCustomer(id string) Customer {
 service := *c.customerService
 return service.getCustomer(id)
}
func (c customerController) getOrders(customerId string) []Order {
 service := *c.customerService
 return service.getOrders(customerId)
}