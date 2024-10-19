package toolkit

import (
	"fmt"

	"github.com/o98k-ok/lazy/v2/alfred"
)

func ShowByAlfred(phase string, content string, err error) {
	if err != nil {
		alfred.Log("%v %v", phase, err.Error())
		return
	}
	fmt.Println(content)
}
