package caesarcipher

type PlainText []byte

func GetPlainTextFromString(str string) PlainText {
	return []byte(str)
}

func (pt PlainText) GetStringFromPlainText() string {
	return string(pt)
}

func (pt PlainText) Encrypt(key Key) CipherText {
	length := len(pt)
	cip := make(CipherText, length)
	for i := 0; i < length; i++ {
		cip[i] = (pt)[i] + key.GetByte()
	}
	return cip
}
