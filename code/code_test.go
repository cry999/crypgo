package code

import "testing"

func TestKey_Encrypt(t *testing.T) {
	tests := map[string]struct {
		key     *Key
		plain   string
		want    string
		wantErr bool
	}{
		"encrypted only words on code bool": {
			key:   New(map[string]string{"A": "1", "B": "2"}),
			plain: "A is B, B is C",
			want:  "1 is 2, 2 is C",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.key.Encrypt([]byte(tt.plain))
			if (err != nil) != tt.wantErr {
				t.Errorf("Key.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if string(got) != tt.want {
				t.Errorf("Key.Encrypt() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestKey_Decrypt(t *testing.T) {
	tests := map[string]struct {
		key       *Key
		encrypted string
		want      string
		wantErr   bool
	}{
		"encrypted only words on code bool": {
			key:       New(map[string]string{"A": "1", "B": "2"}),
			encrypted: "1 is 2, 2 is C",
			want:      "A is B, B is C",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.key.Decrypt([]byte(tt.encrypted))
			if (err != nil) != tt.wantErr {
				t.Errorf("Key.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if string(got) != tt.want {
				t.Errorf("Key.Decrypt() = %s, want %s", got, tt.want)
			}
		})
	}
}
