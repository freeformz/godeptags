package foo

import (
	"fmt"

	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("Foo")
}
