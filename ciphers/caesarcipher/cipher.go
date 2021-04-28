package caesarcipher

type CipherText []byte

func GetCipherTextFromString(str string) CipherText {
	return []byte(str)
}

func (cip CipherText) GetStringFromCipherText() string {
	return string(cip)
}

func (cip CipherText) Decrypt(key Key) PlainText {
	length := len(cip)
	pt := make(PlainText, length)
	for i := 0; i < length; i++ {
		pt[i] = (cip)[i] - key.GetByte()
	}
	return pt
}
