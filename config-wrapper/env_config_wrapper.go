package config_wrapper

import "os"

// GetEnvValue retrives value from config by key. Returns empty value if variable not present.
func GetEnvValue(key ConfigKey) ConfigValue {
	return ConfigValue(os.Getenv(string(key)))
}

// LookupEnvValue retrives value from config by key. Returns empty value and false if variable not present.
func LookupEnvValue(key ConfigKey) (ConfigValue, bool) {
	configValStr, ok := os.LookupEnv(string(key))
	return ConfigValue(configValStr), ok
}

type envConfigWrapper struct{}

// GetValue retrives value from config by key. Returns empty value if variable not present.
func (w envConfigWrapper) GetValue(key ConfigKey) ConfigValue {
	return GetEnvValue(key)
}

// LookupValue retrives value from config by key. Returns empty value and false if variable not present.
func (w envConfigWrapper) LookupValue(key ConfigKey) (ConfigValue, bool) {
	return LookupEnvValue(key)
}

// NewEnvConfigWrapper returns a new instance of env config wrapper.
func NewEnvConfigWrapper() ConfigWrapper {
	return &envConfigWrapper{}
}
