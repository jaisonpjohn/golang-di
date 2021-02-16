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
}


func bootstrap(){
	bootstrapContext = map[string]interface{}{}
	register((*orderDaoImpl)(nil), "orderDao")
	register((*customerDaoImpl)(nil), "customerDao")
	register((*customerServiceImpl)(nil), "customerService")
	register((*customerController)(nil), "customerController")
}

func register(candidate interface{}, name string) {
	t := reflect.TypeOf(candidate).Elem()
	val := reflect.New(t)
	for i := 0; i < t.NumField(); i++ {
		f := reflect.Indirect(val).Field(i)
		ft := t.Field(i)
		customName, ok := ft.Tag.Lookup("inject")
		var name string
		if ok {
			if customName != "" {
				name = customName
			} else {
				name = ft.Type.Name()
			}
			if val, ok := bootstrapContext[name]; ok {
				f.Set(reflect.ValueOf(val))
			} else{
				panic("Bootstrap: Couldnt find: "+name)
			}

		}
	}
	bootstrapContext[name] = val.Interface()
}

func SetRoutes(router *mux.Router){
	customerController := bootstrapContext["customerController"].(*customerController)
	customerRouter := router.PathPrefix("/customers").Subrouter()
	customerController.Register(customerRouter)
}