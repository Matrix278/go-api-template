package random

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/go-openapi/strfmt"
	uuid "github.com/nu7hatch/gouuid"
)

func String(stringLength int) string {
	bytes := make([]byte, stringLength)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)[:stringLength]
}

func UUID4() strfmt.UUID4 {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	}

	return strfmt.UUID4(uuid.String())
}
