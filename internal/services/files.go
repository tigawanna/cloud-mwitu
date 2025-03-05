package services

import "os"

func SaveFile(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)	
	return err
}
