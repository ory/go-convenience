package mapx

import "errors"

var ErrKeyDoesNotExist = errors.New("key is not present in map")
var ErrKeyCanNotBeTypeAsserted = errors.New("key could not be type asserted")

func GetString(values map[interface{}]interface{}, key interface{}) (string, error) {
	if v, ok := values[key]; !ok {
		return "", ErrKeyDoesNotExist
	} else if sv, ok := v.(string); !ok {
		return "", ErrKeyCanNotBeTypeAsserted
	} else {
		return sv, nil
	}
}

func GetStringDefault(values map[interface{}]interface{}, key interface{}, defaultValue string) (string) {
	if s, err := GetString(values, key); err == nil {
		return s
	}
	return defaultValue
}
