package vernamcipher

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Key []byte

func (k Key) String() string {
	return string(k)
}

func (k Key) Validate() error {
	keyLength := len(k)
	for i := 0; i < keyLength; i++ {
		if k[i] < 32 || k[i] > 255 {
			fmt.Println(k[i])
			return errors.New("invalid key: cannot be used with vernam cipher")
		}
	}
	return nil
}

func GenerateRandomKey(keyLength int) Key {
	// 32 to 255
	key := make(Key, keyLength)

	i := int(0)
	for i < keyLength {
		rand.Seed(time.Now().UnixNano())
		key[i] = byte(rand.Uint32())
		if key[i] < 32 {
			continue
		}
		i++
	}
	return key
}
