package main

import (
	"fmt"
	"go-aora-api/internal/database"
	"go-aora-api/internal/server"
	"go-aora-api/internal/utils"
)

func main() {
	fmt.Println("Running go aora api.")

	database.Connect();

	server := server.NewServer()

	/** Print the server address before starting */
	fmt.Println("Server address:", server.Addr)

	/** validator indianize */
	utils.InitValidator()

	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}