package json

import (
	"go-api-template/pkg/logger"

	json "github.com/goccy/go-json"
)

func Encode(input interface{}) ([]byte, error) {
	bodyBytes, err := json.Marshal(input)
	if err != nil {
		logger.Errorf("unable to marshal data: %v", err)

		return nil, err
	}

	return bodyBytes, nil
}

func Decode(input []byte, output interface{}) error {
	if err := json.Unmarshal(input, output); err != nil {
		logger.Errorf("unable to unmarshal data: %v", err)

		return err
	}

	return nil
}
