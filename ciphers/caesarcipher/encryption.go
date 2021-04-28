package caesarcipher

func Encrypt(str string, n byte) string {
	plaintxt := GetPlainTextFromString(str)
	cip := plaintxt.Encrypt(Key(n))
	return cip.GetStringFromCipherText()
}
