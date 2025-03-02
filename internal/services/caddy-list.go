package services

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// ParseCaddyfile parses a Caddyfile into a slice of CaddyService structs
func ListCaddyServices(partialName string) ([]CaddyService, error) {
	caddyfile, err := os.ReadFile("/etc/caddy/Caddyfile")
	if err != nil {
		return nil, fmt.Errorf("failed to read Caddyfile: %v", err)
	}
	caddyfileContent := string(caddyfile)
	services := []CaddyService{}
	blockMap := FindCaddyBlocks(caddyfileContent, partialName)
	caddyfileContentLines := strings.Split(caddyfileContent, "\n")
	for blockName, indices := range blockMap {
		caddyfileContentBlock := caddyfileContentLines[indices[0]:indices[1]]
		parsedBlock := ParseCaddyBlock(caddyfileContentBlock)
		services = append(services, CaddyService{
			Domain: blockName,
			Content: strings.Join(caddyfileContentBlock, "\n"),
			Block: parsedBlock,
		})
	}
	return services, nil
}

type CaddyService struct {
	Domain  string              `json:"domain"`
	Content string              `json:"content"`
	Block   []CaddyServiceBlock `json:"block"`
}

func FindCaddyBlocks(input, partialName string) map[string][2]int {
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

type CaddyServiceBlock struct {
	Path   []string `json:"path"`
	Value  []string `json:"value"`
}

func ParseCaddyBlock(lines []string) []CaddyServiceBlock {

	blocks := []CaddyServiceBlock{}

	var currentPath []string
	currentDepth := 0
	var currentBlock *CaddyServiceBlock

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Skip empty lines and comments
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		// Check if line opens a new block
		if strings.HasSuffix(trimmedLine, "{") {
			// Extract block path without the '{'
			blockPath := strings.TrimSuffix(trimmedLine, "{")
			blockPath = strings.TrimSpace(blockPath)

			// Add to current path
			currentPath = append(currentPath, blockPath)
			currentDepth++

			// Create new block
			currentBlock = &CaddyServiceBlock{
				Path:  append([]string{}, currentPath...), // Make a copy of the path slice
				Value: []string{},
			}
			blocks = append(blocks, *currentBlock)

		} else if trimmedLine == "}" {
			// Close the current block
			if len(currentPath) > 0 {
				currentPath = currentPath[:len(currentPath)-1]
				currentDepth--
			}
		} else {
			// This is a value line (like header_up directives)
			// Find the most recent block at the current depth

			// Add the value to the block
			for i := len(blocks) - 1; i >= 0; i-- {
				if reflect.DeepEqual(blocks[i].Path, currentPath) {
					blocks[i].Value = append(blocks[i].Value, trimmedLine)
					break
				}
			}
		}
	}

	return blocks
}

// type RequestBody struct {
// 	MaxSize string `json:"max_size"`
// }

// type ReverseProxy struct {
// 	Address   string           `json:"address"`
// 	Transport TransportConfig  `json:"transport"`
// 	Headers   []HeaderUpConfig `json:"headers"`
// }

// type TransportConfig struct {
// 	Protocol    string `json:"protocol"`
// 	ReadTimeout string `json:"read_timeout"`
// }

// type HeaderUpConfig struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }
