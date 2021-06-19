package vernamcipher

import "testing"

func TestEncrypt(t *testing.T) {
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
			name: "Encrypt String to Plaintext",
			args: args{
				s:   "MOMOTARO",
				key: Key{86, 99, 92, 102, 111, 67, 97, 87},
			},
			want:    CipherText{163, 178, 169, 181, 195, 132, 179, 166}.String(),
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.s, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
