package uuid

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	src := []byte("TESTU5C494A35458D22000135C448")
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	fmt.Printf("%s\n", dst)
	fmt.Printf("%d\n", len(dst))
	fmt.Printf("%d\n", len("TESTU5C494A35458D22000135C448"))

	str := "Hello from ADMFactory.com"
	hx := hex.EncodeToString([]byte(str))
	fmt.Println("String to Hex Golang example")
	fmt.Println()
	fmt.Println(str + " ==> " + hx)
	fmt.Println(len(hx))

	str = "TESTU5C494A35458D22000135C448"
	hx = hex.EncodeToString([]byte(str))
	fmt.Println("String to Hex Golang example")
	fmt.Println()
	fmt.Println(str + " ==> " + hx)
	fmt.Println(len(hx))

}
