package sipherer

import (
	"encoding/base64"
	"errors"
	"fmt"
)

var ErrEmptySecret = errors.New("secret is empty")
var ErrDecodingFailed = errors.New("decoding failed")
var ErrDecryptLengthEmpty = errors.New("decrypt length empty")

func Cipher(rawString, secret string) (string, error) {

	encryptedBytes, err := process([]byte(rawString), []byte(secret))
	if err != nil {
		return "", fmt.Errorf("encrypt : %w", err)
	}

	return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

func Decipher(cipheredString, secret string) (string, error) {

	cipherBytes, err := base64.StdEncoding.DecodeString(cipheredString)
	if err != nil {
		return "", fmt.Errorf("decrypt : %w", ErrDecodingFailed)
	}

	decryptedBytes, err := process(cipherBytes, []byte(secret))
	if err != nil {
		return "", fmt.Errorf("decrypt : %w", err)
	}

	if len(decryptedBytes) == 0 {
		return "", fmt.Errorf("decrypt : %w", ErrDecryptLengthEmpty)
	}
	return string(decryptedBytes), nil
}

func process(input, secret []byte) ([]byte, error) {
	if len(secret) == 0 {
		return nil, ErrEmptySecret
	}

	for idx, val := range input {
		input[idx] = val ^ secret[idx%len(secret)]
	}

	return input, nil
}
