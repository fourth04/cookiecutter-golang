package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

func Config() Provider {
	return defaultConfig
}

func ConfigPtr() *viper.Viper {
	return defaultConfig
}

func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

func init() {
	defaultConfig = readViperConfig("{{cookiecutter.app_name|upper}}")
}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	// global defaults
	{% if cookiecutter.use_logrus_logging == "y" %}
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")
	{% endif %}

	return v
}

// ReloadConfigFromFlagSets reload config of a configProvider from a flagSet, if the cfgFileKey is not "", it will load cfgFileKey flag in the command flags
func ReloadConfigFromFlagSet(v *viper.Viper, flagSet *pflag.FlagSet, cfgFileKey string) {
	if err := v.BindPFlags(flagSet); err != nil {
		fmt.Printf("err:%s\n", err)
		os.Exit(1)
	}

	cfgFile := v.GetString(cfgFileKey)
	ReloadConfigFromCfgFile(v, cfgFile)
}

// ReloadConfigFromCfgFile reloads config of a configProvider from a config filepath
func ReloadConfigFromCfgFile(v *viper.Viper, cfgFile string) {
	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
		v.SetConfigType(filepath.Ext(cfgFile)[1:])

		if err := v.ReadInConfig(); err != nil {
			fmt.Printf("err:%s\n", err)
			os.Exit(1)
		}
	}
}

func ConfigFileUsed() string {
	return defaultConfig.ConfigFileUsed()
}

func Get(key string) interface{} {
	return defaultConfig.Get(key)
}

func GetBool(key string) bool {
	return defaultConfig.GetBool(key)
}

func GetDuration(key string) time.Duration {
	return defaultConfig.GetDuration(key)
}

func GetFloat64(key string) float64 {
	return defaultConfig.GetFloat64(key)
}

func GetInt(key string) int {
	return defaultConfig.GetInt(key)
}

func GetInt64(key string) int64 {
	return defaultConfig.GetInt64(key)
}

func GetSizeInBytes(key string) uint {
	return defaultConfig.GetSizeInBytes(key)
}

func GetString(key string) string {
	return defaultConfig.GetString(key)
}

func GetStringMap(key string) map[string]interface{} {
	return defaultConfig.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return defaultConfig.GetStringMapString(key)
}

func GetStringMapStringSlice(key string) map[string][]string {
	return defaultConfig.GetStringMapStringSlice(key)
}

func GetStringSlice(key string) []string {
	return defaultConfig.GetStringSlice(key)
}

func GetTime(key string) time.Time {
	return defaultConfig.GetTime(key)
}

func InConfig(key string) bool {
	return defaultConfig.InConfig(key)
}

func IsSet(key string) bool {
	return defaultConfig.IsSet(key)
}
