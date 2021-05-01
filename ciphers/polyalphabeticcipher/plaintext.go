package polyalphabeticcipher

type PlainText []byte

func GetPlainTextFromString(str string) PlainText {
	return PlainText(str)
}

func (pt PlainText) String() string {
	return string(pt)
}

func (pt PlainText) Encrypt(key Key) (CipherText, error) {
	err := key.Validate()
	if err != nil {
		return CipherText{}, nil
	}

	length := len(pt)
	cip := make(CipherText, length)

	for i, v := range pt {
		cursor := byte(i) % key.blockSize

		conversionRule := key.conversionRule[cursor]

		if conversionRule > (126 - v) {
			conversionRule = conversionRule - (126 - v)
			conversionRule = conversionRule % (126 - 31)
			if conversionRule == 0 {
				cip[i] = 126
			} else {
				cip[i] = 31 + conversionRule
			}

		} else {
			cip[i] = pt[i] + conversionRule
		}
	}

	return cip, nil
}
