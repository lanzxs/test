package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var p *int
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
