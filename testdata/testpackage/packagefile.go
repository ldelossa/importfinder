package testpackage

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Mock struct {
	m    sync.Mutex
	desc string
}

func (m *Mock) Print() {
	fmt.Println(m.desc)
}

func (m *Mock) JSON() {
	b, _ := json.Marshal(m.desc)
	fmt.Println(string(b))
}
