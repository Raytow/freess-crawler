package test_methods

import (
	"fmt"
	"log"
	"testing"
)

func TestTcpPing(t *testing.T) {
	a, b, d := tcpPing("34.218.41.53", "443")

	log.Println("avg delay: ", a)
	log.Println("rate: ", b)
	for e := d.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, ", ")
	}
	fmt.Println()
}

func TestGooglePing(t *testing.T) {
	a, b, d := googlePing("34.218.41.53", "443")

	log.Println("avg delay: ", a)
	log.Println("rate: ", b)
	for e := d.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, ", ")
	}
	fmt.Println()
}