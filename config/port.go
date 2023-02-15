package config

import (
	log "github.com/siruspen/logrus"
)

type Type string

const (
	// Local local config
	Local Type = "local"
)

// Factory returns the requested config repo
func Factory(c Type) Repository {
	switch Type(c) {
	case Local:
		return NewLocalRepository()
	default:
		log.Fatalf("Unknown config repository: %s", c)
	}

	return nil
}

// Repository port
type Repository interface {
	Get(key string) (interface{}, error)
	GetInt(key string) (int64, error)
	GetFloat(key string) (float64, error)
	GetString(key string) (string, error)
	GetBool(key string) (bool, error)
	GetStringSlice(key string) ([]string, error)
}
