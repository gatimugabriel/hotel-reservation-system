package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
)

var cryptKey []byte

func init() {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	key, err := loadCryptKey()
	if err != nil {
		fmt.Println("Error loading AES key (CRYPT KEY):", err)
		return
	}

	cryptKey = key
}

// EncryptPayload -> encrypts an object/string
// @returns: encrypted string
func EncryptPayload(payload map[string]interface{}) (string, error) {
	// convert payload to JSON
	plaintext, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	//create AES cipher block
	block, err := aes.NewCipher(cryptKey)
	if err != nil {
		return "", err
	}

	// Generate a new random IV (Initialization Vector)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	// Create a stream cipher using CFB (Cipher Feedback Mode)
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// Encode the ciphertext to base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil

}

// DecryptPayload : decrypts an encrypted string
// @returns:  original payload (dynamic)
func DecryptPayload(encrypted string) (map[string]interface{}, error) {
	// Decode the base64 string
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	// Create AES cipher block
	block, err := aes.NewCipher(cryptKey)
	if err != nil {
		return nil, err
	}

	// Get the IV from the ciphertext
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Create a stream cipher using CFB mode
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt the ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	// Unmarshal the decrypted JSON into a map
	var payload map[string]interface{}
	err = json.Unmarshal(ciphertext, &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// Load the Crypt key from environment variable (base64 encoded)
func loadCryptKey() ([]byte, error) {
	base64Key := os.Getenv("CRYPT_KEY")
	if base64Key == "" {
		return nil, fmt.Errorf("AES_SECRET_KEY not set in environment")
	}

	// Decode base64 key
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return nil, err
	}

	// Validate key size (should be 16, 24, or 32 bytes)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("invalid AES key size: %d", len(key))
	}

	return key, nil
}