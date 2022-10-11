// json
// json.Marshal()
// json.Unmarshal()

package main

import "encoding/json"

// Encode object(map/struct) -> json
func Encode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Decode json -> object(map/struct)
func Decode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}
