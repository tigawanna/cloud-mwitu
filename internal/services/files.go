package services

import "os"

func SaveFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)	
	return err
}
