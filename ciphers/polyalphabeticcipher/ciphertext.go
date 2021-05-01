package polyalphabeticcipher

type CipherText []byte

func GetCipherTextFromString(str string) CipherText {
	return CipherText(str)
}

func (cip CipherText) String() string {
	return string(cip)
}

func (cip CipherText) Decrypt(key Key) (PlainText, error) {
	err := key.Validate()
	if err != nil {
		return PlainText{}, nil
	}

	length := len(cip)
	pt := make(PlainText, length)

	for i, v := range cip {
		cursor := byte(i) % key.blockSize

		conversionRule := key.conversionRule[cursor]

		if conversionRule > (v - 32) {
			conversionRule = conversionRule - (v - 32)
			conversionRule = conversionRule % (126 - 31)
			if conversionRule == 0 {
				pt[i] = 32
			} else {
				pt[i] = 127 - conversionRule
			}

		} else {
			pt[i] = cip[i] - conversionRule
		}
	}

	return pt, nil
}
