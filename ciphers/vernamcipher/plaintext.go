package vernamcipher

import (
	"fmt"
)

type PlainText []byte

func (pt PlainText) String() string {
	// not an optimised way to do it
	// will work for now
	str := ""
	for _, v := range pt {
		str += fmt.Sprintf("%c", v)
	}
	return string(pt)
}

func NewPlaintext(s string) PlainText {
	//plainText := make([]byte, len(s))
	plainText := []byte(s)
	return plainText
}

func (pt PlainText) Validate() error {
	length := len(pt)
	for i := 0; i < length; i++ {
		if pt[i] < 32 || pt[i] > 255 {
			return fmt.Errorf("invalid plaintext: invalid character at position %d", i+1)
		}

	}
	return nil
}

// Vernam's Algorithm
func (p PlainText) Encrypt(k Key) (CipherText, error) {
	// validate the key and cipherText
	err := p.Validate()
	if err != nil {
		return CipherText{}, err
	}
	err = k.Validate()
	if err != nil {
		return CipherText{}, err
	}
	// verify if the length of the key and the plaintext are equal
	if len(p) != len(k) {
		return CipherText{}, fmt.Errorf("invalid key length: key lenght:%d is not matching the plaintext length:%d", len(k), len(p))
	}
	length := len(p)
	// add the key to the plaintext
	// adding can result in bigger number than 255 so create another slice of uint16
	cipherText := make(CipherText, length)
	for i := 0; i < length; i++ {
		temp := uint16(p[i]) + uint16(k[i])
		temp = temp % 255
		cipherText[i] = byte(temp)
	}

	return (cipherText), nil
}
