package main

import (
	"execrise-system/router"
)

func main() {
	r := router.Router()
	r.Run() // listen and serve on
}
