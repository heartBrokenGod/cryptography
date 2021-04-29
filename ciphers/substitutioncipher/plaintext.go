package substitutioncipher

import "errors"

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
		return nil, errors.New("invalid encryption key")
	}

	length := len(pt)
	cip := make(CipherText, length)

	for i, v := range pt {
		cip[i] = key[v]
	}

	return cip, nil

}
