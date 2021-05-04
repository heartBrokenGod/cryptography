package transpositioncipher

import "errors"

type Key struct {
	blockSize        byte
	substitutionRule []byte
}

func (k Key) Validate() error {
	if k.blockSize < 2 {
		return errors.New("invalid key block size is 1 or less")
	}

	subsRuleSliceLength := len(k.substitutionRule)

	if k.substitutionRule == nil || subsRuleSliceLength != int(k.blockSize) {
		return errors.New("invalid key substitution rule array is nil or its length is not equal to block size")
	}

	for i := 0; i < subsRuleSliceLength; i++ {
		if k.substitutionRule[i] > byte(subsRuleSliceLength) || k.substitutionRule[i] < 1 {
			return errors.New("invalid key substitution rule contains index index out of bounds")
		}
		for j := i + 1; j < subsRuleSliceLength; j++ {
			if k.substitutionRule[i] == k.substitutionRule[j] {
				return errors.New("invalid key substitution rule index value is repeated")
			}
		}
	}

	return nil
}

func (k Key) GetBlockSize() byte {
	return k.blockSize
}
func (k Key) GetSubstutionRule() []byte {
	return k.substitutionRule
}

func GetNewKey(blockSize byte, substitutionRule []byte) (Key, error) {
	key := Key{blockSize: blockSize, substitutionRule: substitutionRule}
	err := key.Validate()
	if err != nil {
		return Key{}, err
	}
	return key, nil
}
