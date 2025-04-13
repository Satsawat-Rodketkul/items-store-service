package utils

import (
	"github.com/google/uuid"
)

func GenerateTransactionId() string {
	return "TXN-" + uuid.NewString()[:8]
}
