package vault

import "encoding/json"

type Secret struct {
	data map[string]interface{}
}

func (s Secret) GetString(name string, defaultVal string) string {
	value := s.data[name]
	if value == nil {
		return defaultVal
	}
	return value.(string)
}

func (s Secret) GetBool(name string, defaultVal bool) bool {
	value := s.data[name]
	if value == nil {
		return defaultVal
	}
	return value.(bool)
}

func (s Secret) GetInt(name string, defaultVal int64) (int64, error) {
	value := s.data[name]
	if value == nil {
		return defaultVal, nil
	}
	return value.(json.Number).Int64()
}

func (s Secret) GetFloat(name string, defaultVal float64) (float64, error) {
	value := s.data[name]
	if value == nil {
		return defaultVal, nil
	}
	return value.(json.Number).Float64()
}
