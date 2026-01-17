package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Versions map[string]Version `json:"versions"`
}

type Version struct {
	DLLink          string `json:"dllink"`
	DestinationPath string `json:"destinationPath"`
}

func defaultConfig() Config {
	return Config{
		Versions: map[string]Version{
			"ClassicMop": {
				DLLink:          "https://example.com/classicmop",
				DestinationPath: "",
			},
			"Midnight": {
				DLLink:          "https://example.com/midnight",
				DestinationPath: "",
			},
		},
	}
}

func CreateDefaultIfMissing(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	cfg := defaultConfig()

	data, err := json.MarshalIndent(cfg, "", "   ")
	if err != nil {
		return err
	}
	data = append(data, '\n')

	return os.WriteFile(path, data, 0644)
}

func Load(path)
