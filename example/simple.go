package example

import (
	"github.com/go-pax/mongorm/v0"
	"log"
)

func main() {
	if _, err := mongorm.Connect("mongodb://localhost:27017"); err != nil {
		log.Fatal(err)
	}

}
