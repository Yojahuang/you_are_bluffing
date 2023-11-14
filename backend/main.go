package main

import (
	Router "you_are_bluffing/backend/router"
)

func main() {
	router := Router.RouterFactory()
	router.Run("localhost:8080")
}
