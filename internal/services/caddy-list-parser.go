package services

import (
	"fmt"
	"os"
	"strings"
)



// ParseCaddyfile parses a Caddyfile into a slice of CaddyService structs
func ListCaddyServicesWithFields(partialName string) ([]CaddyService, error) {
	caddyfile, err := os.ReadFile("/etc/caddy/Caddyfile")
	if err != nil {
		return nil, fmt.Errorf("failed to read Caddyfile: %v", err)
	}
	caddyfileContent := string(caddyfile)
	services := []CaddyService{}
	blockMap := FindCaddyBlocks(caddyfileContent,partialName)
	caddyfileContentLines := strings.Split(caddyfileContent, "\n")
	for key, indices := range blockMap {
		caddyfileContentBlock := caddyfileContentLines[indices[0]:indices[1]]
		parsedBlock := ParseCaddyBlock(key, caddyfileContentBlock)
		services = append(services, parsedBlock)
	}
	return services, nil
}



type CaddyParsedService struct {
	Domain  string `json:"domain"`
	Content string `json:"content"`
}


func FindCaddyBlocksWithFields(input, partialName string) map[string][2]int {
	lines := strings.Split(input, "\n")
	blockMap := make(map[string][2]int)
	var currentKey string
	var startLine int
	currentDepth := 0

	for lineIdx, line := range lines {
		keyExtracted := false
		for _, char := range line {
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			if partialName != "" && !strings.Contains(line, partialName) {
				continue
			}
			if char == '{' {
				if currentDepth == 0 && !keyExtracted {
					// Extract the key from the line up to the first '{'
					keyPart := line[:strings.Index(line, "{")]
					currentKey = strings.TrimSpace(keyPart)
					startLine = lineIdx
					keyExtracted = true
				}
				currentDepth++
			} else if char == '}' {
				currentDepth--
				if currentDepth == 0 && currentKey != "" {
					blockMap[currentKey] = [2]int{startLine, lineIdx}
					currentKey = ""
				}
			}
		}
	}

	// Print the results
	// for key, indices := range blockMap {
	// 	fmt.Printf("%s: [%d, %d]\n", key, indices[0], indices[1])
	// }
	return blockMap
}

func ParseCaddyBlockWithFields(blockName string, blockslice []string) CaddyService {
	caddyService := CaddyService{}
	caddyContent := ""
	caddyService.Domain = blockName
	for _, line := range blockslice {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fmt.Println(blockName, "==", line)
		caddyContent += line + "\n"
		caddyService.Content = caddyContent

	}

	return caddyService
}


