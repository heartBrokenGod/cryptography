package transpositioncipher

import "testing"

type encrypttest struct {
	pt, cip string
	key     Key
}

func TestEncrypt(t *testing.T) {
	key, _ := GetNewKey(4, []byte{2, 4, 1, 3})
	encrypttests := []encrypttest{
		{"MOMOTARO", "OOMMAOTR", key},
	}
	for _, encrypttest := range encrypttests {
		result, err := Encrypt(encrypttest.pt, encrypttest.key)
		if err != nil {
			t.Errorf("error : %v", err)
		}
		if result != encrypttest.cip {
			t.Errorf("Result %q not equal to expected %q", result, encrypttest.cip)
		}
	}
}
