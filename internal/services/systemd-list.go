package services


import (
	"bufio"
	"bytes"
	"log"
	"os/exec"
	"strings"
)

type SystemDService struct {
    Name        string
    Unit        string
    ActiveState string
    SubState    string
    LoadState   string
    Path        string
}

func GetSystemDServices(partialName string) []SystemDService {
    cmd := exec.Command("systemctl", "list-units", "--type=service", "--state=active")
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("Could not run systemctl: %v", err)
    }
    scanner := bufio.NewScanner(bytes.NewReader(output))
    scanner.Split(bufio.ScanLines)
    services := []SystemDService{}
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
        service := SystemDService{
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
        service.Path = "/etc/systemd/system/" + service.Name + ".service"
        services = append(services, service)
    }
    return services
}
