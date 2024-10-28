package models

import (
	"os"
	"errors"
)

func InitConfig() error {
	var (
		err error
	)
	NewConfig.Apikey = os.Getenv("API_KEY")
	if NewConfig.Apikey == "" {
		err = errors.New("API_KEY is not set")
		return err
	}
	NewConfig.BaseURL = os.Getenv("BASE_URL")
	if NewConfig.BaseURL == "" {
		err = errors.New("BASE_URL is not set")
		return err
	}
	
	return nil
}