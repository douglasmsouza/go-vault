package vault

import (
	"encoding/json"
	"os"
	"strconv"
)

type Secret struct {
	data map[string]interface{}
}

func NewEmptySecret() *Secret {
	return &Secret{data: make(map[string]interface{})}
}

func (s Secret) GetString(name, defaultVal string) string {
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

func (s Secret) GetInt(name string, defaultVal int) (int, error) {
	value := s.data[name]
	if value == nil {
		return defaultVal, nil
	}
	i, err := value.(json.Number).Int64()
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func (s Secret) GetFloat(name string, defaultVal float64) (float64, error) {
	value := s.data[name]
	if value == nil {
		return defaultVal, nil
	}
	return value.(json.Number).Float64()
}

func (s Secret) GetStringOrEnv(name, env, defaultVal string) string {
	value := s.data[name]
	if value == nil {
		value = os.Getenv(env)
		if value == "" {
			return defaultVal
		}
	}
	return value.(string)
}

func (s Secret) GetBoolOrEnv(name, env string, defaultVal bool) bool {
	value := s.data[name]
	if value == nil {
		value = os.Getenv(env)
		if value == "" {
			return defaultVal
		}

		b, _ := strconv.ParseBool(value.(string))
		return b
	}
	return value.(bool)
}

func (s Secret) GetIntOrEnv(name, env string, defaultVal int) (int, error) {
	value := s.data[name]
	if value == nil {
		value = os.Getenv(env)
		if value == "" {
			return defaultVal, nil
		}

		i, err := strconv.Atoi(value.(string))
		if err != nil {
			return 0, err
		}

		return i, nil
	}

	i, err := value.(json.Number).Int64()
	if err != nil {
		return 0, err
	}

	return int(i), nil
}

func (s Secret) GetFloatOrEnv(name, env string, defaultVal float64) (float64, error) {
	value := s.data[name]
	if value == nil {
		value = os.Getenv(env)
		if value == "" {
			return defaultVal, nil
		}

		i, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			return 0, err
		}

		return i, nil
	}
	return value.(json.Number).Float64()
}
