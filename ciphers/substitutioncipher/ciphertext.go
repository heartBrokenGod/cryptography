package substitutioncipher

import "errors"

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
		return nil, errors.New("invalid encryption key")
	}

	length := len(cip)

	pt := make(PlainText, length)

	for i, v := range cip {

		for k, val := range key {
			if v == val {
				pt[i] = k
				break
			}
		}
	}

	return pt, nil
}
