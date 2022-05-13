package caesar

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct{ key byte }
	tests := map[string]struct {
		args args
		want Key
	}{
		"ok": {args{3}, 3},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := New(tt.args.key); got != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKey_Encrypt(t *testing.T) {
	type args struct{ plain []byte }
	type want struct {
		encrypted []byte
		err       bool
	}
	tests := map[string]struct {
		key  Key
		args args
		want want
	}{
		"ok":                     {3, args{[]byte("Akademia")}, want{[]byte("Dndghpld"), false}},
		"ng/include not alpabet": {3, args{[]byte("!AKADEMIA")}, want{nil, true}},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.key.Encrypt(tt.args.plain)
			if (err != nil) != tt.want.err {
				t.Errorf("Key.Encrypt() error = %v", err)
			}
			if got, want := string(got), string(tt.want.encrypted); got != want {
				t.Errorf("Key.Encrypt() = %v, want %v", got, want)
			}
		})
	}
}

func TestKey_Decrypt(t *testing.T) {
	type args struct{ encrypted []byte }
	type want struct {
		decrypted []byte
		err       bool
	}
	tests := map[string]struct {
		key  Key
		args args
		want want
	}{
		"ok":                     {3, args{[]byte("Dndghpld")}, want{[]byte("Akademia"), false}},
		"ng/include not alpabet": {3, args{[]byte("!AKADEMIA")}, want{nil, true}},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.key.Decrypt(tt.args.encrypted)
			if (err != nil) != tt.want.err {
				t.Errorf("Key.Decrypt() error = %v", err)
			}
			if got, want := string(got), string(tt.want.decrypted); got != want {
				t.Errorf("Key.Decrypt() = %v, want %v", got, want)
			}
		})
	}
}

func Example() {
	const s = "YNRJKQNJXQNPJFSFWWTB"
	got, err := New(5).Decrypt([]byte(s))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", got)

	// Output:
	// TIMEFLIESLIKEANARROW
}
