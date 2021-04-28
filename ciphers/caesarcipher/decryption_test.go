package caesarcipher

import "testing"

type decrypttest struct {
	pt, cip string
	n       byte
}

func TestDecrypt(t *testing.T) {
	decrypttests := []decrypttest{
		{"abc", "bcd", 1},
		{"mno", "nop", 1},
		{"stu", "tuv", 1},
		{"abc", "cde", 2},
	}
	for _, decrypttest := range decrypttests {
		if result := Decrypt(decrypttest.cip, decrypttest.n); result != decrypttest.pt {
			t.Errorf("Result %q not equal to expected %q", result, decrypttest.cip)
		}
	}
}
