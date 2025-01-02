package app

import "github.com/joho/godotenv"

func init() {
	godotenv.Load(".env")
}

func StartApp() error {
	return nil
}

func FirstStart() error {
	return nil
}

func Shutdown() {}
