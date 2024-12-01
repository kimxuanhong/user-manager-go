package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Scheme   string `yaml:"scheme"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
}

var (
	instanceConfig *Config
	configOnce     sync.Once
)

func InitConfig() *Config {
	configOnce.Do(func() {
		instanceConfig = &Config{}
		instanceConfig.load()
	})
	return instanceConfig
}

func loadFile(filename string, config *Config) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", filename, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Cannot close application.yml file!!!")
		}
	}(file)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return fmt.Errorf("failed to decode YAML file %s: %v", filename, err)
	}
	return nil
}

func (r *Config) load() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Không thể lấy thư mục hiện tại: %v", err)
	}

	resourceDir := filepath.Join(currentDir, "resources")
	defaultConfigPath := filepath.Join(resourceDir, "application.yml")
	if err := loadFile(defaultConfigPath, r); err != nil {
		log.Fatalf("Không thể load file application.yml: %v", err)
	}

	if env := os.Getenv("APP_ENV"); env != "" {
		envConfigPath := filepath.Join(resourceDir, fmt.Sprintf("application-%s.yml", env))
		if err := loadFile(envConfigPath, r); err != nil {
			log.Printf("Không thể load file cấu hình môi trường %s: %v", envConfigPath, err)
		}
	}
}
