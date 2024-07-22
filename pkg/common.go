package pkg

import "github.com/google/uuid"

func GetRandomUUIDValue() string {
	requestID := uuid.New()
	return requestID.String()
}
