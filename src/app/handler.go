package app

import (
	"fmt"
	"log"
	"redir/src/component/job"
	"redir/src/core"
	"strings"
)

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

// --job --new --title example --detail example --start 192831223 --end 18239421
func handleCreateJob(parts []string) {
	params := map[string]string{}
	for i, arg := range parts {
		if strings.Contains(arg, "--") {
			params[strings.ReplaceAll(arg, "--", "")] = parts[i+1]
		}
	}
	err := core.AddToRamMemory(params, &job.Jobs, &job.Job{})
	if err != nil {
		log.Println("Create Job Error:", err)
		return
	}
	log.Printf("Job Created: Title=%s, Detail=%s, will be reminded at,", params["title"], params["detail"])
	fmt.Println(job.Jobs)
}

// --job --list
func handleListJob() {
	core.ReadCoreModel(job.Jobs, essentialDirs[0]+listFile)
}
