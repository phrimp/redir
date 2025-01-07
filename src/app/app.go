package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"redir/src/pkg"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var (
	appName       = "Redir"
	essentialDirs = []string{"data", "log"}
	tmpStorage    = map[string]*map[uuid.UUID]interface{}{}
	pidFile       = "redir.pid"
	commandFile   = "command.txt"
	listFile      = "data.txt"
	data_path     = []*string{&pidFile, &commandFile}
	current_os    string
	location, _   = time.LoadLocation("Asia/Ho_Chi_Minh")
)

func init() {
	godotenv.Load(".env")
	if !pkg.CreateMultipleDirs(essentialDirs) {
		log.Fatal("Essential Directories failed to initialize")
	}
	current_os = runtime.GOOS
	if current_os == "windows" {
		for _, file_name := range data_path {
			for i := range essentialDirs {
				essentialDirs[i] = essentialDirs[i] + "\\"
			}
			*file_name = essentialDirs[0] + *file_name
		}
	} else {
		for _, file_name := range data_path {
			for i := range essentialDirs {
				essentialDirs[i] = essentialDirs[i] + "/"
			}
			*file_name = essentialDirs[0] + *file_name
		}
	}
}

func FirstStart() error {
	// Goroutine for log file (new log file each 24h)
	go func() {
		for {
			now := time.Now().In(location)
			log_name := essentialDirs[1] + "log_server_at_" + now.Format("2006-01-02_15-04") + ".log"
			logDir := filepath.Dir(log_name)
			err := pkg.CreateDir(logDir)
			if err != nil {
				log.Fatal("Log file failed to initialize:", err)
			}
			f, err := os.OpenFile(log_name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatal("error opening file:", err)
			}
			log.SetOutput(f)
			time.Sleep(24 * time.Hour)
			f.Close()
		}
	}()
	return nil
}

func StartApp() error {
	if err := enforceSingleInstance(); err != nil {
		return err
	}
	defer removePIDFile()

	if err := FirstStart(); err != nil {
		log.Fatalf("Error init first start: %v", err)
	}
	log.Println("Application running as a service...")

	// Monitor command file for new commands
	go monitorCommands()

	select {} // Keep the service running
}

func Shutdown() {
	log.Println("Cleaning up resources...")
	pkg.ClearDirectory(essentialDirs[0])
}

// Ensure only a single instance of the application runs
func enforceSingleInstance() error {
	if _, err := os.Stat(pidFile); err == nil {
		return fmt.Errorf("application already running")
	}

	pid := []byte(fmt.Sprintf("%d", os.Getpid()))
	if err := os.WriteFile(pidFile, pid, 0644); err != nil {
		return fmt.Errorf("failed to create PID file: %w", err)
	}

	return nil
}

func removePIDFile() {
	os.Remove(pidFile)
}

func removeCommandFile() {
	os.Remove(commandFile)
}

// Handle commands sent to the running instance
func HandleCommand(args []string) error {
	command := strings.Join(args, " ")
	return writeCommand(command)
}

// Write a command to the command file
func writeCommand(command string) error {
	return os.WriteFile(commandFile, []byte(command), 0644)
}

// Monitor the command file for new commands
func monitorCommands() {
	var lastCommand string
	for {
		time.Sleep(1 * time.Second)

		content, err := os.ReadFile(commandFile)
		if err != nil {
			continue // File might not exist yet
		}

		command := string(content)
		if command != lastCommand { // Detect a new command
			lastCommand = command
			handleCommand(command)
		}
	}
}

// Process commands dynamically
func handleCommand(command string) {
	log.Printf("Processing command: %s\n", command)

	parts := strings.Split(command, " ")
	if len(parts) < 1 {
		log.Println("Invalid command")
		return
	}

	switch parts[0] {
	case "--createuser":
		handleCreateUser(parts)
	case "--job":
		switch parts[1] {
		case "--new":
			handleCreateJob(parts)
		case "--list":
			handleListJob()
		}
	default:
		log.Printf("Unknown command: %s\n", command)
	}
}
