package main

import (
	"example.com/m/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
