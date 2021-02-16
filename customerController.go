package main

import (
 "encoding/json"
 "github.com/gorilla/mux"
 "net/http"
)

type customerController struct {
 CService customerService `inject:""`
}

func (c customerController) Register(customerRouter *mux.Router) {
 customerRouter.Path("/{id}").HandlerFunc(c.getCustomer).Methods("GET")
}
func (c customerController) getCustomer(w http.ResponseWriter, r *http.Request) {
 id := mux.Vars(r)["id"]
 customer := c.CService.getCustomer(id)
 jsonValue, _ := json.Marshal(customer)
 w.Header().Set("Content-Type", "application/json")
 w.WriteHeader(http.StatusOK)
 _, _ = w.Write(jsonValue)

}
func (c customerController) getOrders(customerId string) []Order {
 return c.CService.getOrders(customerId)
}
