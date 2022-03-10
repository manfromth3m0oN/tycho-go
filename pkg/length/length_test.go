package length_test

import (
	"log"
	"testing"

	"github.com/manfromth3m0oN/tycho-go/pkg/length"
)

func Test_ReadLength(t *testing.T) {
	buf := make([]byte, 3)
	buf[0] = 0xFF
	buf[1] = 0xFF
	buf[2] = 0x7F

	log.Printf("Buffer: %08b", buf)

	result, err := length.ReadLength(buf)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	log.Printf("Result: %d %016b", result, result)
	if result != 2097151 {
		t.Fail()
	}
}
