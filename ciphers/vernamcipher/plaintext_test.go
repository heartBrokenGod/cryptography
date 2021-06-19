package vernamcipher

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlainText_Encrypt(t *testing.T) {
	type args struct {
		k Key
	}
	tests := []struct {
		name    string
		p       PlainText
		args    args
		want    CipherText
		wantErr bool
	}{
		{
			name: "MOMOTARO Encryption",
			p:    PlainText([]byte("MOMOTARO")),
			args: args{
				k: Key{86, 99, 92, 102, 111, 67, 97, 87}, //
			},
			want:    []byte{163, 178, 169, 181, 195, 132, 179, 166},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.Encrypt(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlainText.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				for _, v := range got {
					fmt.Printf("%d,", v)
				}
				fmt.Println("")
				t.Errorf("PlainText.Encrypt() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestPlainText_String(t *testing.T) {
	tests := []struct {
		name string
		pt   PlainText
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pt.String(); got != tt.want {
				t.Errorf("PlainText.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPlaintext(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want PlainText
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlaintext(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlaintext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlainText_Validate(t *testing.T) {
	tests := []struct {
		name    string
		pt      PlainText
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pt.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("PlainText.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
