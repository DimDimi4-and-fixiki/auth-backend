package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Database Database `yaml:"database"`
	Service  Service  `yaml:"service"`
	Vault    Vault
	Secrets  map[string]string `yaml:"secrets"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Service struct {
	Debug       bool   `yaml:"debug"`
	SwaggerPass string `yaml:"swagger_pass"`
}

type Vault struct {
	Host  string `env:"VAULT_HOST"`
	Port  string `env:"VAULT_PORT"`
	Token string `env:"VAULT_TOKEN"`
	Path  string `env:"VAULT_PATH"`
}

func LoadConfig(env string) (*Config, error) {

	filename := fmt.Sprintf("config/%s/values.yaml", env)

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	vaultConfig, err := readVaultConfig()
	if err != nil {
		return nil, err
	}
	config.Vault = *vaultConfig
	return &config, nil
}

func readVaultConfig() (*Vault, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	vault := Vault{}

	vault.Host = os.Getenv("VAULT_HOST")
	vault.Port = os.Getenv("VAULT_PORT")
	vault.Token = os.Getenv("VAULT_TOKEN")
	vault.Path = os.Getenv("VAULT_PATH")

	return &vault, nil
}
