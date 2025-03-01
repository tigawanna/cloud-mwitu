package services

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type RunningSystemDService struct {
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	ActiveState string `json:"activeState"`
	SubState    string `json:"subState"`
	LoadState   string `json:"loadState"`
	Path        string `json:"path"`
	ModifiedAt  string `json:"modifiedAt"`
}
type SystemDService struct {
	Name        string `json:"name"`
    Contents     string `json:"contents"`
	Path        string `json:"path"`
	ModifiedAt  string `json:"modifiedAt"`
}




func GetSystemDServiceFiles(partialName string,etc bool) []SystemDService {
	services := []SystemDService{}
	//    /etc/systemd/system/
	var dir string
	if etc {
		dir = "/etc"
	} else {
		dir = "/lib"
	}
	files, err := os.ReadDir(dir + "/systemd/system/")
	if err != nil {
		log.Fatalf("Could not read directory: %v", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".service") {
			if partialName != "" && !strings.Contains(file.Name(), partialName) {
				continue
			}
            contents, err := os.ReadFile(dir + "/systemd/system/" + file.Name())
            if err != nil {
                log.Printf("Could not read file: %v", err)
            }
			service := SystemDService{
				Name: file.Name(),
                Path: "/etc/systemd/system/" + file.Name(),
                Contents: string(contents),
			}
          
          if fileInfo, err := os.Stat(service.Path); err == nil {
			// Note: CreatedAt is not available in Linux directly
			// Using ModTime for both since creation time isn't reliably stored
			// fmt.Println("file mod time", fileInfo.ModTime())
			if fileInfo.ModTime().IsZero() {
				service.ModifiedAt = "Not available"
			} else {
				service.ModifiedAt = fileInfo.ModTime().Format("Jan 02 2006 15:04:05")
			}
		}
			services = append(services, service)
		}
	}

	return services
}
func GetRunningSystemDServices(partialName string) []RunningSystemDService {
	cmd := exec.Command("systemctl", "list-units", "--type=service", "--state=active")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Could not run systemctl: %v", err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(output))
	scanner.Split(bufio.ScanLines)
	services := []RunningSystemDService{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "UNIT") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		if partialName != "" && !strings.Contains(fields[0], partialName) {
			continue
		}
		service := RunningSystemDService{
			Name:        fields[0],
			Unit:        fields[1],
			ActiveState: fields[2],
		}
		if len(fields) > 3 {
			service.SubState = fields[3]
		}
		if len(fields) > 4 {
			service.LoadState = fields[4]
		}
		service.Path = "/etc/systemd/system/" + service.Name

		// Get file info for timestamps
		if fileInfo, err := os.Stat(service.Path); err == nil {
			// Note: CreatedAt is not available in Linux directly
			// Using ModTime for both since creation time isn't reliably stored
			fmt.Println("file mod time", fileInfo.ModTime())
			if fileInfo.ModTime().IsZero() {
				service.ModifiedAt = "Not available"
			} else {
				service.ModifiedAt = fileInfo.ModTime().Format("Jan 02 2006 15:04:05")
			}
		}

		services = append(services, service)
	}
	return services
}
