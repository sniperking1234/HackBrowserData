//go:build linux

package crypto

func Chromium(key, encryptPass []byte) ([]byte, error) {
	if len(encryptPass) < 3 {
		return nil, errPasswordIsEmpty
	}

	chromeIV := []byte{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32}
	return aes128CBCDecrypt(key, chromeIV, encryptPass[3:])
}

func DPAPI(data []byte) ([]byte, error) {
	return nil, nil
}
