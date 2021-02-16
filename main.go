package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
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
	var o orderDaoImpl
	register(o, "orderDao")

	var c customerDaoImpl
	register(c, "customerDao")
	var d customerServiceImpl
	register(d, "customerService")
	var contlr customerController
	register(contlr, "customerController")
}

func register(candidate interface{}, name string) {
	t := reflect.TypeOf(candidate)
	vava := reflect.New(t)
	fmt.Println("---t.NumField()---",t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := reflect.Indirect(vava).Field(i)
		ft := t.Field(i)
		customName, ok := ft.Tag.Lookup("inject")
		var name string
		if ok {
			if customName != "" {
				name = customName
			} else {
				name = ft.Type.Name()
			}
			fmt.Println("f",f,"---ft:",ft, "--FtTag:",name)
			if val, ok := bootstrapContext[name]; ok {
				f.Set(reflect.ValueOf(val))
			} else{
				panic("Bootstrap: Couldnt find: "+name)
			}

		}
	}
	bootstrapContext[name] = vava.Interface()
}

func SetRoutes(router *mux.Router){
	customerController := bootstrapContext["customerController"].(*customerController)
	customerRouter := router.PathPrefix("/customers").Subrouter()
	customerController.Register(customerRouter)
}