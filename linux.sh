#./linux.sh /path/to/your/project
#!/bin/bash

# Function to print messages
print() {
  echo -e "\033[1;32m$1\033[0m"
}

print "Starting the Go project setup script for Linux..."

# Install Go using a package manager
if ! command -v go &>/dev/null; then
  if command -v apt &>/dev/null; then
    print "Using apt to install Golang..."
    sudo apt update && sudo apt install -y golang
  elif command -v yum &>/dev/null; then
    print "Using yum to install Golang..."
    sudo yum install -y golang
  elif command -v brew &>/dev/null; then
    print "Using Homebrew to install Golang..."
    brew install go
  else
    print "No supported package manager found! Please install Golang manually."
    exit 1
  fi
else
  print "Golang is already installed. Skipping installation."
fi

# Verify Golang installation
if ! command -v go &>/dev/null; then
  echo "Failed to install Golang. Exiting."
  exit 1
fi

print "Golang installed successfully. Version: $(go version)"

# Check for an existing Go project
PROJECT_DIR=$1

if [ -z "$PROJECT_DIR" ]; then
  print "No project directory specified. Usage: ./linux.sh /path/to/your/project"
  exit 1
fi

if [ ! -f "$PROJECT_DIR/go.mod" ]; then
  print "No 'go.mod' file found in $PROJECT_DIR. Ensure this is a valid Go project."
  exit 1
fi

print "Navigating to project directory: $PROJECT_DIR"
cd "$PROJECT_DIR" || exit

# Download all project dependencies
print "Downloading all dependencies for the Go project..."
go mod download

if [ $? -eq 0 ]; then
  print "Dependencies downloaded successfully!"
else
  echo "Failed to download dependencies. Please check for errors above."
  exit 1
fi

# Verify the project dependencies
print "Verifying dependencies..."
go mod tidy

print "Go project dependencies setup complete. You're ready to code!"
