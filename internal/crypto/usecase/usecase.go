package usecase

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"

	cryptomodels "konverter/internal/crypto/models"

	"golang.org/x/crypto/hkdf"
)

// deriveKey derives a 32-byte AES key from the secret using HKDF or SHA-256
func deriveKey(secret, salt, ctxInfo string) ([]byte, error) {
	secretBytes := []byte(secret)

	// If both salt and ctxInfo are provided, use HKDF
	if salt != "" && ctxInfo != "" {
		saltBytes := []byte(salt)
		info := []byte(ctxInfo)

		// Use HKDF with SHA-256 to derive a 32-byte key
		hash := sha256.New
		hkdf := hkdf.New(hash, secretBytes, saltBytes, info)

		key := make([]byte, 32) // AES-256 key size
		if _, err := io.ReadFull(hkdf, key); err != nil {
			return nil, errors.New("failed to derive key using HKDF: " + err.Error())
		}

		return key, nil
	}

	// Otherwise, use SHA-256 hash of the secret
	hash := sha256.Sum256(secretBytes)
	return hash[:], nil
}

// Encrypt encrypts the plaintext using AES-GCM
func Encrypt(req cryptomodels.EncryptRequest) (cryptomodels.EncryptResponse, error) {
	if err := req.Validate(); err != nil {
		return cryptomodels.EncryptResponse{}, err
	}

	// Derive the encryption key
	key, err := deriveKey(req.Secret, req.Salt, req.CtxInfo)
	if err != nil {
		return cryptomodels.EncryptResponse{}, err
	}

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return cryptomodels.EncryptResponse{}, errors.New("failed to create AES cipher: " + err.Error())
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return cryptomodels.EncryptResponse{}, errors.New("failed to create GCM: " + err.Error())
	}

	// Generate random nonce (12 bytes for GCM)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return cryptomodels.EncryptResponse{}, errors.New("failed to generate nonce: " + err.Error())
	}

	// Encrypt the plaintext
	plaintext := []byte(req.Text)
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// Encode as Base64
	encryptedText := base64.StdEncoding.EncodeToString(ciphertext)

	return cryptomodels.EncryptResponse{
		EncryptedText: encryptedText,
	}, nil
}

// Decrypt decrypts the Base64 encoded ciphertext using AES-GCM
func Decrypt(req cryptomodels.DecryptRequest) (cryptomodels.DecryptResponse, error) {
	if err := req.Validate(); err != nil {
		return cryptomodels.DecryptResponse{}, err
	}

	// Decode Base64 input
	ciphertext, err := base64.URLEncoding.DecodeString(req.Text)
	if err != nil {
		return cryptomodels.DecryptResponse{}, errors.New("failed to decode Base64: " + err.Error())
	}

	// Derive the decryption key (same as encryption)
	key, err := deriveKey(req.Secret, req.Salt, req.CtxInfo)
	if err != nil {
		return cryptomodels.DecryptResponse{}, err
	}

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return cryptomodels.DecryptResponse{}, errors.New("failed to create AES cipher: " + err.Error())
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return cryptomodels.DecryptResponse{}, errors.New("failed to create GCM: " + err.Error())
	}

	// Check if ciphertext is long enough to contain nonce
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return cryptomodels.DecryptResponse{}, errors.New("ciphertext too short")
	}

	// Extract nonce and ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt the ciphertext
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return cryptomodels.DecryptResponse{}, errors.New("failed to decrypt: " + err.Error())
	}

	return cryptomodels.DecryptResponse{
		DecryptedText: string(plaintext),
	}, nil
}

