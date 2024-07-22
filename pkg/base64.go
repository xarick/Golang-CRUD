package pkg

import (
	"encoding/base64"
	"errors"
	"strings"
)

func BasicAuthLogPass(basicAuth string) (login string, password string, err error) {
	base64Data := basicAuth[6:]
	decodedByte, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		err = errors.New("error receiving login and password")
		return
	}

	decodedString := string(decodedByte)
	logAndPas := strings.Split(decodedString, ":")

	login = logAndPas[0]
	password = logAndPas[1]
	return
}
