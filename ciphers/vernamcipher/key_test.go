package vernamcipher

import (
	"reflect"
	"testing"
)

func TestKey(t *testing.T) {

	// check if the key is generating and is valid
	t.Run("check if the key is generating and valid", func(t *testing.T) {
		key := GenerateRandomKey(20)

		if key != nil && len(key) != 20 {
			t.Errorf("either key is nil or not of specified length")
		}

		if key.Validate() != nil {
			t.Errorf("key validation failed: generated key is not valid")
		}

	})
}

func TestKey_String(t *testing.T) {
	tests := []struct {
		name string
		k    Key
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("Key.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_Validate(t *testing.T) {
	tests := []struct {
		name    string
		k       Key
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.k.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Key.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateRandomKey(t *testing.T) {
	type args struct {
		keyLength int
	}
	tests := []struct {
		name string
		args args
		want Key
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomKey(tt.args.keyLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
