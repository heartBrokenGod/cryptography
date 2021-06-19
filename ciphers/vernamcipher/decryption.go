package vernamcipher

func Decrypt(s string, key Key) (string, error) {
	ct := NewCipherText(s)
	pt, err := ct.Decrypt(key)
	if err != nil {
		return "", err
	}
	return pt.String(), nil
}
