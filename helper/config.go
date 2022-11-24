package helper

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/viper"
)

// GetConfig : Get Config
func GetConfig(key string) string {
	return viper.GetString(key)
}

func GetConfigInt(key string) int {
	return viper.GetInt(key)
}

// InitConfig : init Config
func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("fatal error config file: %s", err)
	}
}

func GetDecodedFireBaseKey() ([]byte, error) {

	fireBaseAuthKey := GetConfig("FUNDCLICK.FIREBASE_AUTH_KEY")

	decodedKey, err := base64.StdEncoding.DecodeString(fireBaseAuthKey)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}
