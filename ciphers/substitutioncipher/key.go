package substitutioncipher

import (
	"errors"
	"math/rand"
	"time"
)

type Key map[byte]byte

func (k Key) ByteMap() map[byte]byte {
	return map[byte]byte(k)
}

func GenerateRandomKey() (Key, error) {
	var byteArr [94]byte
	var byteInitialiser byte

	byteInitialiser = 32
	for i := range byteArr {
		byteArr[i] = byteInitialiser
		byteInitialiser++
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(byteArr), func(i, j int) { byteArr[i], byteArr[j] = byteArr[j], byteArr[i] })

	key := make(Key)

	byteInitialiser = 32
	for _, val := range byteArr {
		key[byteInitialiser] = val
		byteInitialiser++
	}
	err := key.Validate()

	if err != nil {
		return nil, err
	}
	return key, nil
}

func GetKeyFromMap(m map[byte]byte) (Key, error) {
	key := Key(m)
	err := key.Validate()

	if err != nil {
		return nil, err
	}
	return key, nil
}

func (key Key) Validate() error {
	for k, v := range key {
		if k > 126 && k < 32 && v > 126 && v < 32 {
			return errors.New("invalid key unable to use with substitution cipher")
		}
	}
	return nil
}
