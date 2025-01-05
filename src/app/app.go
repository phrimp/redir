package app

import (
	"redir/src/pkg"

	"github.com/joho/godotenv"
)

// APPLICATION INITIALIZATION
func init() {
	godotenv.Load("local.env")
	app_name = "Redir"
	if !pkg.CreateMultipleDir(essential_dir) {
		panic("Essential Directories failed to initialize")
	}
}

var (
	app_name      string
	essential_dir = []string{"data", "log"}
)

func StartApp() error {
	return nil
}

func FirstStart() error {
	return nil
}

func Shutdown() {}
