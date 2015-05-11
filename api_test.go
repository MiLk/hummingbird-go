package hummingbird

import (
	"fmt"

	"github.com/milk/hummingbird"
)

// You can use the API by following this example
func ExampleNewAPI() {
	api := hummingbird.NewAPI()
	_, results := api.Search("One Piece")
	fmt.Print("%+v", results)
}
