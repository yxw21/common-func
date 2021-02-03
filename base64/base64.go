package base64

import (
	"encoding/base64"
	"strconv"
)

func Encode(content interface{}) string {
	switch c := content.(type) {
	case string:
		return base64.StdEncoding.EncodeToString([]byte(c))
	case int:
		return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(c)))
	case []byte:
		return base64.StdEncoding.EncodeToString(c)
	default:
		return ""
	}
}

func Decode(content string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return "", err
	}
	return string(decodeBytes), nil
}
