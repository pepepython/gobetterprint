package gobetterprint

import (
	"fmt"
	"math/rand"
	"time"
)

func rdm(n, m int) int {
	return rand.Intn(m-n) + n
}

func Print(interface{}) {
	fmt.Println("Processing....,,,..")
	time.Sleep(4 * time.Second)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(".")
	for i := 0; i < 20; i++ {
		time.Sleep(time.Millisecond * time.Duration(rdm(100, 700)))
		for p, r := 0, rdm(4, 20); p < r; p++ {
			fmt.Printf(".")
		}
		fmt.Println(".")
	}
	panic("error memory leak out of memory")
}
