package main

import (
	"flag"
	"fmt"
	"github.com/grjotwani/sojernHW/versioncheck"
	"log"
)

func main() {
	version1 := flag.String(
		"version1",
		"0.1",
		"version1 string")

	version2 := flag.String(
		"version2",
		"0.09",
		"version2 string")

	flag.Parse()

	if *version1 == "" || *version2 == "" {
		log.Fatal("Please provide version1 and version2")
	}
	v1 := versioncheck.New(*version1)
	fmt.Printf("output : %d", v1.Compare(*version2))
}
