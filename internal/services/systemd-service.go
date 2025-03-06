package services




type SystemDFileServiceImpl struct{}

func NewSystemDFileService() *SystemDFileServiceImpl {
    return &SystemDFileServiceImpl{}
}

type SystemDFileService interface{
	GetSystemDServiceFiles(partialName string, libDir bool) ([]SystemDService, error)
	GetRunningSystemDServices(partialName string) ([]RunningSystemDService, error)
	NewSystemdFileConfig(serviceName, baseDir, execCommand string, libDir bool, opts *ConfigOptions) (SystemdServiceConfig, error)
	UpdateSystemDFile(path string, newService string, libDir bool) (string, error)
} 
