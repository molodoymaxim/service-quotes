package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Парсинг конфигурации из ENV файла в множеств оструктур
func GetConfigsENV(path, name string, cfg []any) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("reading in config: %v", err)
	}

	for _, c := range cfg {
		if err := viper.Unmarshal(&c); err != nil {
			return fmt.Errorf("unmarshal config: %v", err)
		}
	}

	return nil
}
