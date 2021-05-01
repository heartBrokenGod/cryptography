package polyalphabeticcipher

func Decrypt(str string, key Key) (string, error) {
	cip := GetCipherTextFromString(str)
	pt, err := cip.Decrypt(key)
	if err != nil {
		return "", err
	}
	return pt.String(), nil
}
