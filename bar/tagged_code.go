// +build mytag

package bar

import (
	"fmt"

	"github.com/pkg/errors"
)

func init() {
	fmt.Println(errors.New("Hello"))
}
