package config

import "fmt"

func getBool(key string, r Repository) (bool, error) {
	raw, err := r.Get(key)
	if err != nil {
		return false, err
	}

	secret, ok := raw.(bool)
	if !ok {
		return false, fmt.Errorf("unable to cast config with key: %s to string", key)
	}

	return secret, nil
}

func getFloat(key string, r Repository) (float64, error) {
	raw, err := r.Get(key)
	if err != nil {
		return 0, err
	}

	secret, ok := raw.(float64)
	if !ok {
		return 0, fmt.Errorf("unable to cast config with key: %s to float64", key)
	}

	return secret, nil
}
func getInt(key string, r Repository) (int64, error) {
	raw, err := r.Get(key)
	if err != nil {
		return 0, err
	}

	secret, ok := raw.(int)
	if !ok {
		return 0, fmt.Errorf("unable to cast config with key: %s to int", key)
	}

	return int64(secret), nil
}
func getString(key string, r Repository) (string, error) {
	raw, err := r.Get(key)
	if err != nil {
		return "", err
	}
	secret, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("unable to cast config with key: %s to string", key)
	}

	return secret, nil
}
