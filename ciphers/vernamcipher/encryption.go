package vernamcipher

func Encrypt(s string, key Key) (string, error) {
	// get Plaintext from string
	p := NewPlaintext(s)
	ct, err := p.Encrypt(key)
	if err != nil {
		return "", err
	}
	return ct.String(), nil
}
