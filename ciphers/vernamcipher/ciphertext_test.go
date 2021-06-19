package vernamcipher

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCipherText_Decrypt(t *testing.T) {
	type args struct {
		k Key
	}
	tests := []struct {
		name    string
		c       CipherText
		args    args
		want    PlainText
		wantErr bool
	}{
		{
			name: "MOMOTORO Decrypt",
			c:    []byte{163, 178, 169, 181, 195, 132, 179, 166},
			args: args{
				k: Key{86, 99, 92, 102, 111, 67, 97, 87},
			},
			want:    []byte(`MOMOTARO`),
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Decrypt(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("CipherText.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				for _, v := range got {
					fmt.Printf("%c", v)
				}
				fmt.Println("")
				t.Errorf("CipherText.Decrypt() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestCipherText_String(t *testing.T) {
	tests := []struct {
		name string
		ct   CipherText
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ct.String(); got != tt.want {
				t.Errorf("CipherText.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCipherText_Validate(t *testing.T) {
	tests := []struct {
		name    string
		ct      CipherText
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ct.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("CipherText.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
