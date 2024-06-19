package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type MyState int

func (u MyState) MarshalJSON() ([]byte, error) {
	encodedAge := fmt.Sprintf("%d years", u)
	encodedAge = strconv.Quote(encodedAge) //  返回之前用引号将字符串括起来
	return []byte(encodedAge), nil
}

type Lan struct {
	B MyState `json:"b,omitempty" tag:"lll"`
	C string
}

func main() {
	var l MyState = 22
	p := &Lan{l, "lan"}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	s := strings.Builder{}
	s.WriteString("lan")
	s.WriteString(" zhang")
	fmt.Println(s.String())
	i := -1 << 6
	fmt.Println("i 01 100", i)
}
