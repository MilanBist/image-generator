package main

import (
	"github.com/image-generator/config"
	"github.com/image-generator/routes"
)

func main(){
	// obtain the port 
	c := config.Port()

	// set all the routers with all the middlewares
	router := routes.SetRouter()
	// run the router in the obtained port
	port := ":"+c.Port
	router.Run(port)
}


