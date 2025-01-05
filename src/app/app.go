package app

import (
	"fmt"
	"log"
	"os"
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
	data_path     = []*string{&pidFile, &commandFile}
	current_os    string
)

func init() {
	godotenv.Load(".env")
	if !pkg.CreateMultipleDirs(essentialDirs) {
		panic("Essential Directories failed to initialize")
	}
	current_os = runtime.GOOS
	if current_os == "windows" {
		for _, file_name := range data_path {
			*file_name = essentialDirs[0] + "\\" + *file_name
		}
	} else {
		for _, file_name := range data_path {
			*file_name = essentialDirs[0] + "/" + *file_name
		}
	}
}

func StartApp() error {
	if err := enforceSingleInstance(); err != nil {
		return err
	}
	defer removePIDFile()

	log.Println("Application running as a service...")

	// Monitor command file for new commands
	go monitorCommands()

	select {} // Keep the service running
}

func Shutdown() {
	log.Println("Cleaning up resources...")
	removePIDFile()
	removeCommandFile()
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
	default:
		log.Printf("Unknown command: %s\n", command)
	}
}

func handleCreateUser(parts []string) {
	var name, age string
	for i, arg := range parts {
		if arg == "--name" && i+1 < len(parts) {
			name = parts[i+1]
		}
		if arg == "--age" && i+1 < len(parts) {
			age = parts[i+1]
		}
	}

	if name == "" || age == "" {
		log.Println("Error: --name and --age are required for --createuser")
		return
	}

	log.Printf("User created: Name=%s, Age=%s\n", name, age)
}
