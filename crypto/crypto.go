package crypto

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "fmt"
    "io"
)

func Encrypt(data []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    ciphertext := make([]byte, aes.BlockSize+len(data))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return nil, err
    }

    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(ciphertext[aes.BlockSize:], data)

    return ciphertext, nil
}

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    if len(ciphertext) < aes.BlockSize {
        return nil, fmt.Errorf("ciphertext is too short")
    }

    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(ciphertext, ciphertext)

    return ciphertext, nil
}

func ProcessFile(filePath string, key []byte, encryptFlag, decryptFlag bool) ([]byte, error) {
    // Read the file content
    plaintext, err := file.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    switch {
    case encryptFlag:
        return Encrypt(plaintext, key)
    case decryptFlag:
        return Decrypt(plaintext, key)
    default:
        return nil, fmt.Errorf("either -encrypt or -decrypt flag must be provided")
    }
}
