package main

import (
	"fmt"

	"github.com/google/uuid"
	types "github.com/timmo001/letmeknow-types-go"
)

func GenerateUserID(clientType types.ClientType, additional *string) string {
	uuid := uuid.New().String()

	// Remove the dashes from the UUID
	uuid = uuid[:8] + uuid[9:13] + uuid[14:18] + uuid[19:23] + uuid[24:]

	if additional != nil {
		return fmt.Sprintf("%s-%s-%s", clientType, *additional, uuid)
	}

	return fmt.Sprintf("%s-%s", clientType, uuid)
}
