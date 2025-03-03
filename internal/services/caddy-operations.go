package services

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)



func ReadcaddyFile() (string, error) {
	caddyfile, err := os.ReadFile("/etc/caddy/Caddyfile")
	if err != nil {
		return "", fmt.Errorf("failed to read Caddyfile: %v", err)
	}
	caddyfileContent := string(caddyfile)	
	return caddyfileContent, nil
}

// ParseCaddyfile parses a Caddyfile into a slice of CaddyService structs
func ListCaddyServices(partialName string) ([]CaddyService, error) {
	caddyfileContent,err := ReadcaddyFile()
	if err != nil {
		return nil, err
	}
	services := []CaddyService{}
	blockMap := FindCaddyBlocks(caddyfileContent)
	caddyfileContentLines := strings.Split(caddyfileContent, "\n")
	for blockName, indices := range blockMap {
		caddyfileContentBlock := caddyfileContentLines[indices[0]:indices[1]]
		parsedBlock := ParseCaddyBlock(caddyfileContentBlock)
		if strings.Contains(blockName, partialName) {
			services = append(services, CaddyService{
				Domain:  blockName,
				Content: strings.Join(caddyfileContentBlock, "\n"),
				Block:   parsedBlock,
				StartEnd: indices,
			})

		}
	}
	return services, nil
}

type CaddyService struct {
	Domain  string              `json:"domain"`
	Content string              `json:"content"`
	Block   []CaddyServiceBlock `json:"block"`
	StartEnd [2]int             `json:"startEnd"`
}

func FindCaddyBlocks(input string) map[string][2]int {
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

	return blockMap
}

type CaddyServiceBlock struct {
	Path  []string `json:"path"`
	Value []string `json:"value"`
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

type UpdateCaddyResponse struct {
	UpdatedBlock string `json:"updatedBlock"`
	Content string `json:"content"`
	ContentArrayBefore []string `json:"contentArrayBefore"`
	ContentArray []string `json:"contentArray"`
}

func UpdateCaddyFile(domainName, newConfig string) (UpdateCaddyResponse, error) {
	caddyfileContent, err := ReadcaddyFile()
	foundMatch := false
	updateCaddyResponse := UpdateCaddyResponse{}
	if err != nil {
		return updateCaddyResponse, err
	}
	blockMap := FindCaddyBlocks(caddyfileContent)

	caddyfileContentLines := strings.Split(caddyfileContent, "\n")
	newConfigLines := strings.Split(newConfig, "\n")
	for blockName, indices := range blockMap {
		if blockName == domainName {
			// Step 1: Get all lines before the domain block
			beforeDomainBlock := caddyfileContentLines[:indices[0]]
			// Step 2: Get all lines after the domain block
			afterDomainBlock := caddyfileContentLines[indices[1]+1:]
			// Step 3: Combine the new config lines with the lines after the domain block
			newConfigAndAfter := append(newConfigLines, afterDomainBlock...)
			// Step 4: Combine everything: lines before + new config + lines after
			caddyfileContentLines := append(beforeDomainBlock, newConfigAndAfter...)
			updateCaddyResponse.Content = strings.Join(caddyfileContentLines, "\n")
			updateCaddyResponse.ContentArray = caddyfileContentLines
			updateCaddyResponse.ContentArrayBefore = strings.Split(caddyfileContent, "\n")
			updateCaddyResponse.UpdatedBlock = blockName
			foundMatch = true
			return updateCaddyResponse, nil
			// return {strings.Join(caddyfileContentLines, "\n"}), nil
		}
	}
	if(!foundMatch){
		newCaddyfileContentLines := append(caddyfileContentLines, newConfigLines...)
		updateCaddyResponse.Content = strings.Join(newCaddyfileContentLines, "\n")
		updateCaddyResponse.ContentArray = newCaddyfileContentLines
		updateCaddyResponse.ContentArrayBefore = strings.Split(caddyfileContent, "\n")
		updateCaddyResponse.UpdatedBlock = domainName
	}
	return updateCaddyResponse, nil
}
