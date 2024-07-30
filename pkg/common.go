package pkg

import "github.com/google/uuid"

func GetUUID() string {
	requestID := uuid.New()
	return requestID.String()
}
