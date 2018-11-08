package marshal

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MarshalWrap returns the JSON Marshaled string wrapped in a key.
func Wrap(s interface{}, key string) (string, error) {
	if len(key) < 1 {
		return "", errors.New("Key must be at least 1 character")
	}

	j, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("{\"%s\": %s}", key, j), nil
}
