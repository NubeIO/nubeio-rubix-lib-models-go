package dto

import (
	"bytes"
	"encoding/json"
	"strconv"
)

/*
a := "{\"a\": 99999999}"
var b interface{}
_ = json.Unmarshal([]byte(a), &b)
fmt.Println("b", b) // it prints `b map[a:9.9999999e+07]`
*/

func (m *DynamicMap) UnmarshalJSON(data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	var rawMap map[string]interface{}
	if err := decoder.Decode(&rawMap); err != nil {
		return err
	}

	*m = make(map[string]interface{})
	for key, value := range rawMap {
		switch v := value.(type) {
		case json.Number:
			// Attempt to parse the number as an integer
			if intValue, err := v.Int64(); err == nil {
				(*m)[key] = intValue
			} else {
				// If parsing as an integer fails, parse as a float
				floatValue, _ := strconv.ParseFloat(v.String(), 64)
				(*m)[key] = floatValue
			}
		case map[string]interface{}:
			// Recursively process nested maps
			nestedMap := make(DynamicMap)
			if err := nestedMap.populateFromMap(v); err != nil {
				return err
			}
			(*m)[key] = nestedMap
		default:
			(*m)[key] = v
		}
	}

	return nil
}

func (m *DynamicMap) populateFromMap(data map[string]interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(marshalMap(data)))
	decoder.UseNumber()

	return decoder.Decode(m)
}

func marshalMap(data map[string]interface{}) []byte {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return dataBytes
}

type HostConfig struct {
	AppName      string     `json:"app_name,omitempty"`
	Body         DynamicMap `json:"body"`                  // used when writing JSON, YML data
	BodyAsString string     `json:"body_as_string"`        // used when writing string data
	ConfigName   string     `json:"config_name,omitempty"` // config.yml
}

type HostConfigResponse struct {
	FilePath string `json:"file_path,omitempty"`
	Data     []byte `json:"data"`
}
