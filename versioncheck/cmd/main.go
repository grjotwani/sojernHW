package main

import (
	"fmt"
	"github.com/grjotwani/sojernHW/versioncheck"
)

func main() {
	v1 := versioncheck.New("5.0.1.11.3.2")
	fmt.Println(v1.Compare("5.0.1.12"))

	v1 = versioncheck.New("5.0.1.12.0")
	fmt.Println(v1.Compare("5.0.1.12"))
}