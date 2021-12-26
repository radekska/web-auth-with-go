package hmacUtils

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)

func generateKey() []byte {
	var key = []byte{}
	for i := 0; i < 64; i++ {
		key = append(key, byte(i))
	}
	return key
}

func SignMessage(message []byte) ([]byte, error) {
	key := generateKey()
	hmacHash := hmac.New(sha512.New, key)
	_, err := hmacHash.Write(message)
	if err != nil {
		return nil, fmt.Errorf("error while hashing the message: %w", err)
	}
	signature := hmacHash.Sum(nil)
	println(signature)
	return signature, nil
}

func checkSignature(message string, signature []byte) (bool, error) {
	newSignature, err := SignMessage([]byte(message))
	if err != nil {
		return false, fmt.Errorf("error in checkSignature while getting sigunature of the message: %w", err)
	}
	areSame := hmac.Equal(newSignature, signature)
	return areSame, nil

}
