package services

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)



type CaddyFileModelBlock struct {
	Path  []string `json:"path"`
	Value []string `json:"value"`
}


// ParseCaddyBlock parses a Caddyfile block into a slice of CaddyFileModelBlock structs.
//
// A block is defined as a set of lines between two brackets, e.g.:
// {
//     foo
//     bar
// }
//
// The returned slice contains one or more CaddyFileModelBlock structs, each with a
// Path field that contains the path of the block, and a Value field that contains
// the value of the block.
func ParseCaddyBlock(lines []string) []CaddyFileModelBlock {

	blocks := []CaddyFileModelBlock{}

	var currentPath []string
	currentDepth := 0
	var currentBlock *CaddyFileModelBlock

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
			currentBlock = &CaddyFileModelBlock{
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



func ReadcaddyFile(configPath ...string) (string, error) {
	var path string
	if len(configPath) > 0 {
		path = configPath[0]
	} else {
		path = "/etc/caddy/Caddyfile"
	}
	caddyfile, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read Caddyfile: %v", err)
	}
	caddyfileContent := string(caddyfile)	
	return caddyfileContent, nil
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
