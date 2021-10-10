package converter

import "encoding/json"

func ConvertStructToString(value interface{}) (string, error) {
	convert, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(convert), nil
}

