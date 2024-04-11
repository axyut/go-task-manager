package cmd

import "testing"

func TestEncodingAndDecoding(t *testing.T) {
	given := "This is very long normal text"
	want := given

	encoded := EncodeBase64([]byte(given))
	decoded := DecodeBase64(encoded)

	if string(decoded) != want {
		t.Fatal("given: "+given, "want: "+want, "got: "+string(decoded))
	}
}
