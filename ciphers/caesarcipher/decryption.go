package caesarcipher

func Decrypt(str string, n byte) string {
	cip := GetCipherTextFromString(str)
	pt := cip.Decrypt(GetKeyFromByte(n))
	return pt.GetStringFromPlainText()
}
