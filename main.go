package main

import (
	"fmt"
)

func main() {

	result := monitools.avgCPUOverNSeconds(5)
	fmt.Println(result)
}
