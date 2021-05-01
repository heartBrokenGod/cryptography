package polyalphabeticcipher

import "errors"

type Key struct {
	blockSize      byte
	conversionRule []byte
}

func (k Key) Validate() error {
	if k.blockSize < 1 {
		return errors.New("invalid key block size is 0 or less")
	}

	if k.conversionRule == nil || len(k.conversionRule) > int(k.blockSize) {
		return errors.New("invalid key conversion rule array is nil or its length is greater than block size")
	}

	return nil
}

func (k Key) GetBlockSize() byte {
	return k.blockSize
}
func (k Key) GetConversionRule() []byte {
	return k.conversionRule
}

func GetNewKey(blockSize byte, conversionRule []byte) (Key, error) {
	key := Key{blockSize: blockSize, conversionRule: conversionRule}
	err := key.Validate()
	if err != nil {
		return Key{}, err
	}
	return key, nil
}
