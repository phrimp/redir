package app

import (
	"fmt"
	"log"
	"redir/src/component/job"
	"redir/src/core"
	"strconv"
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
	params := map[string]interface{}{}
	for i, arg := range parts {
		switch arg {
		case "--title":
			params["title"] = parts[i+1]
		case "--detail":
			params["detail"] = parts[i+1]
		case "--start":
			tmp_start, err := strconv.Atoi(parts[i+1])
			if err != nil {
				log.Println("Error: Converting --start to int failed:", err)
				return
			}
			params["start"] = int64(tmp_start)
		case "--end":
			tmp_end, err := strconv.Atoi(parts[i+1])
			if err != nil {
				log.Println("Error: Converting --end to int failed:", err)
				return
			}
			params["end"] = int64(tmp_end)
		}
	}
	if params["title"] == nil || params["start"] == nil || params["end"] == nil {
		log.Println("Error: --title ; --start ; --end are required for --job --new")
		return
	}
	err := core.AddToRamMemory(params, &job.Jobs, &job.Job{})
	if err != nil {
		log.Println("Create Job Error:", err)
		return
	}
	log.Printf("Job Created: Title=%s, Detail=%s, will be reminded at,", params["title"], params["detail"])
	fmt.Println(job.Jobs)
}
