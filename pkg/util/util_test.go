package util

import (
	"fmt"
	"testing"

	"github.com/forgoer/openssl"
)

func TestAuth_Login(t *testing.T) {
	code := "5050fd57038a4078196d5a250edf8026a0a0074a0001a36d4273fb93d9016e2878399e81888a4fbcf9814e5f3f2a4f121757de3d1415a4fc9f032930ef98a662eebf093b4c9189b341ada7d6d84f4ac8"
	src := HexToByte(code)
	iv := []byte("e4750b34230b48e1")
	key := []byte("b0891a7f6018e5a76b085e3cb9548edd")

	result, err := openssl.AesCBCDecrypt(src, key, iv, openssl.PKCS7_PADDING)
	fmt.Println(string(result), err)
}
