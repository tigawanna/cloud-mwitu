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
}
type SystemDService struct {
	Name       string `json:"name"`
	Contents   string `json:"contents"`
	Path       string `json:"path"`
	ModifiedAt string `json:"modifiedAt"`
}

func (s *SystemDFileServiceImpl) GetSystemDServiceFiles(partialName string, libDir bool) ([]SystemDService, error) {
	fmt.Println("=== GetSystemDServiceFiles : partialName, libDir === ", partialName, libDir)
	services := []SystemDService{}
	//    /etc/systemd/system/
	var dir string
	if libDir {
		dir = "/lib"
	} else {
		dir = "/etc"
	}
	files, err := os.ReadDir(dir + "/systemd/system/")
	if err != nil {
		log.Printf("Could not read directory: %v", err)
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".service") {
			if partialName != "" && !strings.Contains(file.Name(), partialName) {
				continue
			}
			contents, err := os.ReadFile(dir + "/systemd/system/" + file.Name())
			if err != nil {
				log.Printf("Could not read file: %v", err)
				continue
			}
			service := SystemDService{
				Name:     file.Name(),
				Path:     "/etc/systemd/system/" + file.Name(),
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

	return services,nil
}
// func (s *SystemDFileServiceImpl) GetSystemDServiceFileByName(partialName string, libDir bool) (SystemDService, error) {
// 	services,err := s.GetSystemDServiceFiles(partialName, libDir)
// 	if err != nil {
// 		return SystemDService{}, err
// 	}
// 	fmt.Println("services ==", services)
// 	oneService := SystemDService{}
// 	//    /etc/systemd/system/
// 	if(len(services) == 0){
// 		return oneService, fmt.Errorf("no caddy file item found with name %s", partialName)
// 	}
// 	oneService = services[0]
// 	return oneService,nil
// }
func (s *SystemDFileServiceImpl) GetRunningSystemDServices(partialName string) ([]RunningSystemDService, error) {
	cmd := exec.Command("systemctl", "list-units", "--type=service", "--state=active")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Could not run systemctl: %v", err)
		return nil, err
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

		services = append(services, service)
	}
	return services, nil
}
