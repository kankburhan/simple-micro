package util

import (
	"encoding/json"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	type Data struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Id      string `json:"id"`
	}

	signature := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4OTAiLCJuYW1lIjoiVGVzdCIsImFkZHJlc3MiOiJpbmRvbmVzaWEifQ.jV8x6vNtGXAVOdtwjcM8X8rFzDFUA3XXwfOtWvzLOkg"

	result, err := ValidateJwt(signature, "dGVzdGtsaWtkb2t0ZXJpbmRvbmVzaWE=")

	var payload Data
	jsonData, _ := json.Marshal(result.Claims.(jwt.MapClaims))
	_ = json.Unmarshal(jsonData, &payload)

	assert.Nil(t, err, "should be not error")
	assert.Equal(t, "Test", payload.Name, "Name must be Test")
	assert.Equal(t, "indonesia", payload.Address, "Name must be indonesia")
}
