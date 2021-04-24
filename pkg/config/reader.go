package config

import (
	"os"

	"github.com/facup94/worldcup-predictor/pkg/environment"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Name               string `yaml:"name"`
	Host               string `yaml:"host"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	MaxOpenConnections int    `yaml:"max_open_connections"`
}

type allDBsConfig struct {
	Local      DatabaseConfig `yaml:"local"`
	Staging    DatabaseConfig `yaml:"staging"`
	Production DatabaseConfig `yaml:"production"`
}

func GetDBConfig(env environment.Environment) (*DatabaseConfig, error) {
	// Open config file
	file, err := os.Open("db_config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Create config structure
	var config allDBsConfig

	// Start YAML decoding from file
	if err = d.Decode(&config); err != nil {
		return nil, err
	}

	var cfg DatabaseConfig

	switch env {
	case environment.Local:
		cfg = config.Local
	case environment.Staging:
		cfg = config.Staging
	case environment.Production:
		cfg = config.Production
	}

	cfg.Host = os.Getenv("DB_HOST")
	cfg.Password = os.Getenv("DB_PASSWORD")

	return &cfg, nil
}
