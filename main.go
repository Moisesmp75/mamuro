package main

import (
	"flag"
	"mamuro_mail-api/api"
)

func main() {
	var port int
	// fmt.Println(os.Getenv("port"))
	flag.IntVar(&port, "port", 3000, "Puerto para el servidor")
	flag.Parse()
	// fmt.Println(port)
	api.RunServer(port)
}