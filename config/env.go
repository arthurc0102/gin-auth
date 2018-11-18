package config

import (
	"log"
	"os"
)

func checkEnv(name string, _default ...string) string {
	value := os.Getenv(name)

	if value == "" {
		if len(_default) < 1 {
			log.Fatalf("Environ %s not set\n", name)
		}

		os.Setenv(name, _default[0])
		return _default[0]
	}

	return value
}
