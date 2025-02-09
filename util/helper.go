package util

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
)

func StructToJson(StructToJsonValue any) string {
	jsonData, err := json.Marshal(StructToJsonValue)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error"
	}
	return string(jsonData)
}

func JsonToStruct(JsonToStructValue []byte, StructValue any) error {
	err := json.Unmarshal([]byte(JsonToStructValue), &StructValue)
	return err
}

const charsetSecureRandomString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func SecureRandomString(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		randomByte, err := rand.Int(rand.Reader, big.NewInt(int64(len(charsetSecureRandomString))))
		if err != nil {
			return "", err
		}
		result[i] = charsetSecureRandomString[randomByte.Int64()]
	}
	return string(result), nil
}
