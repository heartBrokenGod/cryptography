package vernamcipher

import "fmt"

type CipherText []byte

func NewCipherText(s string) CipherText {
	lenght := len(s)
	cipherText := []byte(s)
	cipherText = cipherText[:lenght]
	return cipherText
}

func (ct CipherText) String() string {
	str := ""
	for _, v := range ct {
		str += fmt.Sprintf("%c", v)
	}
	return str
}

func (ct CipherText) Validate() error {
	length := len(ct)
	for i := 0; i < length; i++ {
		if ct[i] < 32 || ct[i] > 255 {
			return fmt.Errorf("invalid plaintext: invalid character at position %d", i+1)
		}

	}
	return nil
}

func (c CipherText) Decrypt(k Key) (PlainText, error) {
	// validate the key and cipherText
	err := c.Validate()
	if err != nil {
		return PlainText{}, err
	}
	err = k.Validate()
	if err != nil {
		return PlainText{}, err
	}
	// validate the lengths of ciphertext and the plaintext are same
	if len(c) != len(k) {
		return PlainText{}, fmt.Errorf("invalid key length: key lenght:%d is not matching the ciphertext length:%d", len(k), len(c))
	}

	length := len(c)

	plainText := make(PlainText, length)
	for i := 0; i < length; i++ {
		temp := uint16(c[i]) + uint16(256) - uint16(k[i])
		plainText[i] = byte(temp)
	}

	return plainText, nil
}
