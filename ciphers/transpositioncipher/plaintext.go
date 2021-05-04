package transpositioncipher

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
		return CipherText{}, err
	}

	length := len(pt)

	if length%int(key.blockSize) != 0 {
		return CipherText{}, errors.New("invalid plaintext size must be in multiple of key block size")
	}

	cip := make(CipherText, length)
	cursor := 0

	for cursor < length {
		block := pt[cursor : cursor+int(key.blockSize)]

		for j := 0; j < int(key.blockSize); j++ {

			cip[cursor+j] = block[key.substitutionRule[j]-1]

		}

		cursor += int(key.blockSize)
	}

	return cip, nil
}
