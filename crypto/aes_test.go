package crypto

import "testing"

func TestAESEncrypt(t *testing.T) {
	v := "123456"

	cipher, err := AESEncrypt(v)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Cipher: ", cipher)

	decodedValue, err := AESDecrypt(cipher)
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Decoded Value: ", decodedValue)

}
