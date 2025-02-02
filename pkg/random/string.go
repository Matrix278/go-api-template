package random

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/go-openapi/strfmt"
	uuid "github.com/nu7hatch/gouuid"
)

func String(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)[:n]
}

func UUID4() strfmt.UUID4 {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	}

	return strfmt.UUID4(uuid.String())
}
