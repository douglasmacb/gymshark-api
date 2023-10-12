package config

import (
	"fmt"
	"os"
)

const packageSizesEnvPropertyName = "PACKAGE_SIZES"

func PackageSizesFromEnv() (string, error) {
	var env = os.Getenv(packageSizesEnvPropertyName)
	if env == "" {
		return "", fmt.Errorf("missing mandatory environment variable %s", packageSizesEnvPropertyName)
	}

	return env, nil
}
