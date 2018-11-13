package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main() {
	x := math.Abs(float64(1))
	b, _ := json.Marshal(int(x))
	fmt.Println(string(b))
}
