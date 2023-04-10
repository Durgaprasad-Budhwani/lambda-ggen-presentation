package main

import hello_world_golang "github.com/lambda-gqgen-presentation"

func main() {
	router := hello_world_golang.Router()
	router.Run(":8080")
}
