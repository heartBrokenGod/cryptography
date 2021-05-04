package transpositioncipher

import "testing"

type decrypttest struct {
	pt, cip string
	key     Key
}

func TestDecrypt(t *testing.T) {
	key, _ := GetNewKey(4, []byte{2, 4, 1, 3})
	decrypttests := []decrypttest{
		{"MOMOTARO", "OOMMAOTR", key},
	}
	for _, decrypttest := range decrypttests {
		result, err := Decrypt(decrypttest.cip, decrypttest.key)
		if err != nil {
			t.Errorf("error : %v", err)
		}
		if result != decrypttest.pt {
			t.Errorf("Result %q not equal to expected %q", result, decrypttest.pt)
		}
	}
}
