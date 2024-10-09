package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"github.com/ebosetalee/password-service.git/config"
)

var env = config.Env
// Function to encrypt the plaintext
func Encrypt(text []byte) (string, error) {
	key := []byte(env.EncryptionKey)

	// Create a new AES cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Use GCM for authenticated encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate a random nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the data using the AES-GCM
	ciphertext := aesGCM.Seal(nonce, nonce, text, nil)

	// Return the encrypted data as a hex string
	return hex.EncodeToString(ciphertext), nil
}

// Function to decrypt the ciphertext
func Decrypt(ciphertext string, dst interface{}) (error) {

	key := []byte(env.EncryptionKey)

	// Decode the hex string to get the raw ciphertext
	data, err := hex.DecodeString(ciphertext)
	if err != nil {
		return err
	}

	// Create a new AES cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Use GCM for authenticated decryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Extract the nonce from the encrypted data
	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return fmt.Errorf("ciphertext too short")
	}

	nonce := data[:nonceSize] // Use this for the GCM mode
	ciphertextBytes := data[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(plaintext, &dst); err != nil {
		return err
	}
	
	return nil
}
