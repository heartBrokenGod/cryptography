package caesarcipher

import (
	"math/rand"
)

type Key byte

func (k Key) GetByte() byte {
	return byte(k)
}

func GenerateRandomKey() Key {
	key := Key(rand.Intn(5))
	return key
}

func GetKeyFromByte(n byte) Key {
	key := Key(n)
	return key
}
