package vernamcipher

import (
	"testing"
)

func TestDecrypt(t *testing.T) {
	type args struct {
		s   string
		key Key
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Decrypt MOMOTORO",
			args: args{
				s:   string([]byte{163, 178, 169, 181, 195, 132, 179, 166}),
				key: Key{86, 99, 92, 102, 111, 67, 97, 87},
			},
			want:    "MOMOTARO",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.s, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
