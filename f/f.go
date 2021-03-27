package f

import (
	"fmt"
)

type F struct {
	Name string
}

func (f *F) run() {
	fmt.Printf("run f \n")
}
