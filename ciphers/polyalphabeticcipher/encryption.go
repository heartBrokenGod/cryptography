package polyalphabeticcipher

func Encrypt(str string, key Key) (string, error) {
	plaintxt := GetPlainTextFromString(str)
	cip, err := plaintxt.Encrypt(key)
	if err != nil {
		return "", err
	}
	return cip.String(), nil
}
