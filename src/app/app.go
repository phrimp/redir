package app

import (
	"redir/src/core"
	corefunction "redir/src/core/core_function"
	"redir/src/pkg"

	"github.com/google/uuid"
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
	tmp_storage   = map[string]*map[uuid.UUID]core.CoreModel{
		"job": &corefunction.Jobs,
	}
)

func StartApp() error {
	return nil
}

func FirstStart() error {
	return nil
}

func Shutdown() {}
