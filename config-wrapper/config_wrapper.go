package config_wrapper

import (
	"fmt"
	"strconv"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=ConfigWrapper --case=underscore

// ConfigKey defines config key type.
type ConfigKey string

// ConfigValue is returned by ConfigKey
type ConfigValue string

// String converts config value to string.
func (v ConfigValue) String() string {
	return string(v)
}

// Int converts config value to integer. Panics on failure.
func (v ConfigValue) Int() int {
	intValue, err := strconv.Atoi(string(v))
	if err != nil {
		panic(fmt.Errorf("cannot convert config value to type int: %w", err))
	}

	return intValue
}

// ConfigWrapper defines interface for config wrapper for mocking.
type ConfigWrapper interface {
	// GetValue retrives value from config by key. Returns empty value if variable not present.
	GetValue(key ConfigKey) ConfigValue
	// LookupValue retrives value from config by key. Returns empty value and false if variable not present.
	LookupValue(key ConfigKey) (ConfigValue, bool)
}
