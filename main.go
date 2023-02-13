package main

import (
	"github.com/otaxhu/serverX"
)

func main() {
	server := serverX.NewServer(":3000")
	server.Handle("GET", "/", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle("GET", "/products", server.AddMiddleware(GetProducts, Logging()))
	server.Handle("POST", "/products/create", server.AddMiddleware(PostProducts, Logging()))
	server.Handle("PUT", "/products/update", server.AddMiddleware(PutProductByID, Logging()))
	server.Handle("DELETE", "/products/delete", server.AddMiddleware(DeleteProductByID, Logging()))
	server.Listen()
}
