// package-imports/main-package/main.go
package main

import (
	// "main"
	"fmt"
	_ "github.com/lib/pq"

)

func init() {
    fmt.Println("launch initialization")
}

func main() {
    fmt.Println("launch the program !")
}