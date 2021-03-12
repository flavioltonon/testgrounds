package configmap

import "os"

func String(name string, defaultValue ...string) string {
	value := os.Getenv(name)

	if len(value) == 0 && len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return value
}
