package main

import (
	"fmt"
	"go-aora-api/internal/server"
)

func main() {
	fmt.Println("Running go aora api.")

	server := server.NewServer()

	/** Print the server address before starting */
	fmt.Println("Server address:", server.Addr)

	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}