package main

import (
	"fmt"
	"log"
	"net/http"
)
import "github.com/gorilla/mux"

var bootstrapContext map[string]interface{}

func main()  {
	fmt.Println("Server is starting")
	router := mux.NewRouter()
	bootstrap()
	SetRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))
	//fmt.Println("Server is Shutting Down")
}


func bootstrap(){
	bootstrapContext = map[string]interface{}{}
	bootstrapContext["customerDao"] = &customerDaoImpl{} //inject DB pool
	bootstrapContext["orderDao"] = &orderDaoImpl{} //inject DB pool
	bootstrapContext["customerService"] = &customerServiceImpl{
		cDao: bootstrapContext["customerDao"].(customerDao),
		oDao: bootstrapContext["orderDao"].(orderDao),
	}
	bootstrapContext["customerController"] = &customerController{
		customerService: bootstrapContext["customerService"].(customerService),
	}
}

func SetRoutes(router *mux.Router){
	customerController := bootstrapContext["customerController"].(*customerController)
	customerRouter := router.PathPrefix("/customers").Subrouter()
	customerController.Register(customerRouter)
}