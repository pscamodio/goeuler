package euler11

import "testing"
import "fmt"

func SuppressError() {
	_ = fmt.Println
}

func TestEuler11_1(t *testing.T) {
	data, err := LoadGrid("data.txt")
	if err != nil {
		t.Fail()
		fmt.Println(err)
	}
	max := data.MaxQuad()
	fmt.Println(max)
}
