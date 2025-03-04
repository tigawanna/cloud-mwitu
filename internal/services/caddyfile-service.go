package services

import (
	"fmt"
	"strings"
)


type CaddyFileServiceImpl struct {
	caddyFilePath string
}

func NewCaddyFileService(configPath string) *CaddyFileServiceImpl {
    return &CaddyFileServiceImpl{
        caddyFilePath: configPath,
    }
}


// ParseCaddyfile parses a Caddyfile into a slice of CaddyFileModel structs
func (s *CaddyFileServiceImpl) GetCaddyFileItems(partialName string) ([]CaddyFileModel, error) {
	caddyfileContent,err := ReadcaddyFile()
	if err != nil {
		return nil, err
	}
	services := []CaddyFileModel{}
	blockMap := FindCaddyBlocks(caddyfileContent)
	caddyfileContentLines := strings.Split(caddyfileContent, "\n")
	for blockName, indices := range blockMap {
		caddyfileContentBlock := caddyfileContentLines[indices[0]:indices[1]]
		parsedBlock := ParseCaddyBlock(caddyfileContentBlock)
		if strings.Contains(blockName, partialName) {
			services = append(services, CaddyFileModel{
				Domain:  blockName,
				Content: strings.Join(caddyfileContentBlock, "\n"),
				Block:   parsedBlock,
				StartEnd: indices,
			})

		}
	}
	return services, nil
}


// GetCaddyFileItemByName implements the interface method
func (s *CaddyFileServiceImpl) GetCaddyFileItemByName(name string) (CaddyFileModel, error) {
    // Implementation goes here
	caddyFileItems, err := s.GetCaddyFileItems(name)
	if err != nil {
		return CaddyFileModel{}, err
	}
	if len(caddyFileItems) == 0 {
		return CaddyFileModel{}, fmt.Errorf("no caddy file item found with name %s", name)
	}
	caddyFileItem := caddyFileItems[0]
    return caddyFileItem,nil
}


type CaddyFileModel struct {
	Domain  string              `json:"domain"`
	Content string              `json:"content"`
	Block   []CaddyFileModelBlock `json:"block"`
	StartEnd [2]int             `json:"startEnd"`
}

type UpdateCaddyResponse struct {
	UpdatedBlock string `json:"updatedBlock"`
	Content string `json:"content"`
	ContentArrayBefore []string `json:"contentArrayBefore"`
	ContentArray []string `json:"contentArray"`
}

func (s *CaddyFileServiceImpl) UpdateCaddyFile(domainName, newConfig string) (UpdateCaddyResponse, error) {
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


type CaddyFileService interface {
	GetCaddyFileItems(partialName string) ([]CaddyFileModel, error) 
	GetCaddyFileItemByName(name string) (CaddyFileModel, error)
	UpdateCaddyFile(domainName, newConfig string) (UpdateCaddyResponse, error)
}
