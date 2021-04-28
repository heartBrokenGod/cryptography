package caesarcipher

import "testing"

type encrypttest struct {
	pt, cip string
	n       byte
}

func TestEncrypt(t *testing.T) {
	encrypttests := []encrypttest{
		{"abc", "bcd", 1},
		{"mno", "nop", 1},
		{"stu", "tuv", 1},
		{"abc", "cde", 2},
	}
	for _, encrypttest := range encrypttests {
		if result := Encrypt(encrypttest.pt, encrypttest.n); result != encrypttest.cip {
			t.Errorf("Result %q not equal to expected %q", result, encrypttest.cip)
		}
	}
}
