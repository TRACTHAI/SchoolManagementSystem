package decrypt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	keyUsername = "e230944a-e25d-4674-83c7-436f7085086e"
	keyPassword = "LLodeHVjIDV-N13thflXkjWZuu1y4rCo723BGOLQ8RYGAalYETJz5HmsYx5MXwfH3mgTXw93UxtzVfPgzGNYCw"
)

func Detokenize(usernameToken string) (string, error) {
	fortanixAPIURL := "https://sdkms.fortanix.com/crypto/v1/keys/2c197fdf-2db3-4021-8b7e-940630493f6a/decrypt"

	// สร้าง JSON request โดยระบุ "cipher" ที่เป็นค่า "username_token"
	reqBody := fmt.Sprintf(`{"alg": "AES", "mode": "FPE", "cipher": "%s"}`, usernameToken)

	client := &http.Client{}
	req, err := http.NewRequest("POST", fortanixAPIURL, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	// ตั้งค่าการรับรองความถูกต้อง (HTTP Basic Authentication)
	req.SetBasicAuth(keyUsername, keyPassword)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Log ตอบรับ
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// log.Printf("Response from Fortanix API: %s", string(responseBytes))

	// อ่านค่า "plain" จากการเรียก API
	var result struct {
		Plain string `json:"plain"`
	}
	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return "", err
	}

	return result.Plain, nil
}
