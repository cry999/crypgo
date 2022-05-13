package scytale

import (
	"fmt"
	"regexp"
	"testing"
)

func TestKey_Encrypt(t *testing.T) {
	tests := map[string]struct {
		key     Key
		plain   string
		want    string
		wantErr bool
	}{
		"ok/n=2": {New(2), "helloworld", "hweolrllod", false},
		"ok/n=3": {New(3), "helloworld", "holewdlo.lr.", false},
		"ok/n=4": {New(4), "helloworld", "hlodeor.lwl.", false},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := tt.key.Encrypt([]byte(tt.plain))
			if (err != nil) != tt.wantErr {
				t.Errorf("Key.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if ok, err := regexp.MatchString("^"+tt.want+"$", string(got)); err != nil || !ok {
				t.Errorf("Key.Encrypt() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestKey_Decrypt(t *testing.T) {
	tests := map[string]struct {
		key       Key
		encrypted string
		want      string
		wantErr   bool
	}{
		"ok/n=2": {New(2), "hweolrllod", "helloworld", false},
		"ok/n=3": {New(3), "holewdlo.lr.", "helloworld..", false},
		"ok/n=4": {New(4), "hlodeor.lwl.", "helloworld..", false},

		"ng/n=2": {New(2), "hweolrllo", "", true},
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

func Example() {
	const s = "AAMSMTFSCAIYTTOCGSEOMAMCRRELBE"
	key := New(6)
	ans, err := key.Decrypt([]byte(s))
	if err != nil {
		panic(err)
	}
	fmt.Printf("ans: %s\n", ans)
	// Output:
	// ans: AFTERASTORMCOMESACALMIGMBTYSCE
}
