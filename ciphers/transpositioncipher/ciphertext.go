package transpositioncipher

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
		return PlainText{}, err
	}

	length := len(cip)

	if length%int(key.blockSize) != 0 {
		return PlainText{}, errors.New("invalid plaintext size must be in multiple of key block size")
	}
	pt := make(PlainText, length)
	cursor := 0
	for cursor < length {
		block := cip[cursor : cursor+int(key.blockSize)]

		for j := 0; j < int(key.blockSize); j++ {

			for i := 0; i < int(key.blockSize); i++ {

				if j+1 == int(key.substitutionRule[i]) {
					pt[cursor+j] = block[i]
					break
				}

			}

		}

		cursor += int(key.blockSize)
	}

	return pt, nil
}
