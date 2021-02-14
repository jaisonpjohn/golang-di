package main

import (
	"fmt"
	"log"
	"net/http"
)
import "github.com/gorilla/mux"

func main()  {
	fmt.Println("Server is starting")
	router := mux.NewRouter()
	SetRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))
	//fmt.Println("Server is Shutting Down")
}

func SetRoutes(router *mux.Router){
	customerDao := customerDaoImpl{}
	orderDao := orderDaoImpl{}
	cService := customerServiceImpl{
		cDao: &customerDao,
		oDao: &orderDao,
	}
	customerController := customerController{customerService: &cService} // read it from weavedContext
	customerRouter := router.PathPrefix("/customers").Subrouter()
	customerController.Register(customerRouter)
}