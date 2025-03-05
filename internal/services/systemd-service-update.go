package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type SystemdServiceConfig struct {
	Unit    UnitSection
	Service ServiceSection
	Install InstallSection
	Path    string
}

type UnitSection struct {
	Description string
}

type ServiceSection struct {
	Type           string
	User           string
	Group          string
	LimitNOFILE    int
	Restart        string
	RestartSec     string
	StandardOutput string
	StandardError  string
	ExecStart      string
}

type InstallSection struct {
	WantedBy string
}

type ConfigOptions struct {
	Type        string
	User        string
	Group       string
	LimitNOFILE int
	Restart     string
	RestartSec  string
}

const (
	defaultType        = "simple"
	defaultUser        = "root"
	defaultGroup       = "root"
	defaultLimitNOFILE = 4096
	defaultRestart     = "always"
	defaultRestartSec  = "5s"
	defaultWantedBy    = "multi-user.target"
)

// Template for the systemd service file
const serviceTemplate = `[Unit]
Description={{.Unit.Description}}

[Service]
Type={{.Service.Type}}
User={{.Service.User}}
Group={{.Service.Group}}
LimitNOFILE={{.Service.LimitNOFILE}}
Restart={{.Service.Restart}}
RestartSec={{.Service.RestartSec}}
StandardOutput={{.Service.StandardOutput}}
StandardError={{.Service.StandardError}}
ExecStart={{.Service.ExecStart}}

[Install]
WantedBy={{.Install.WantedBy}}
`

func (s *SystemDFileServiceImpl) NewSystemdFileConfig(serviceName, baseDir, execCommand string, libDir bool, opts *ConfigOptions) (SystemdServiceConfig, error) {
	// Default options
	if opts == nil {
		opts = &ConfigOptions{
			Type:        defaultType,
			User:        defaultUser,
			Group:       defaultGroup,
			LimitNOFILE: defaultLimitNOFILE,
			Restart:     defaultRestart,
			RestartSec:  defaultRestartSec,
		}
	}

	// Expand home directory if path starts with ~
	if strings.HasPrefix(baseDir, "~/") {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			baseDir = filepath.Join(homeDir, baseDir[2:])
		}
	}

	// Ensure base directory is absolute
	baseDir, _ = filepath.Abs(baseDir)

	// Build paths
	logPath := filepath.Join(baseDir, "logs", "service.log")
	execPath := filepath.Join(baseDir, execCommand)
	var savePath string
	// /etc is the recommended location but can be overridden by passing ilbDorTrue
	if libDir {
		filepath.Join("/lib/systemd/system", serviceName+".service")
	} else {
		filepath.Join("/etc/systemd/system", serviceName+".service")
	}

	systemdFileConfig := SystemdServiceConfig{
		Unit: UnitSection{
			Description: fmt.Sprintf("%s service", serviceName),
		},
		Service: ServiceSection{
			Type:           opts.Type,
			User:           opts.User,
			Group:          opts.Group,
			LimitNOFILE:    opts.LimitNOFILE,
			Restart:        opts.Restart,
			RestartSec:     opts.RestartSec,
			StandardOutput: "append:" + logPath,
			StandardError:  "append:" + logPath,
			ExecStart:      execPath,
		},
		Install: InstallSection{
			WantedBy: defaultWantedBy,
		},
		Path: savePath,
	}

	systemdFileConfigString, err := systemdFileConfig.ToString()

	if err != nil {
		return systemdFileConfig, err
	}

	err = SaveFile(savePath, systemdFileConfigString)

	if err != nil {
		return systemdFileConfig, err
	}

	return systemdFileConfig, nil
}

func (c SystemdServiceConfig) ToString() (string, error) {
	// Parse the template
	tmpl, err := template.New("systemdService").Parse(serviceTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute the template with the config data
	var sb strings.Builder
	if err := tmpl.Execute(&sb, c); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return sb.String(), nil
}

func (s *SystemDFileServiceImpl) UpdateSystemDFile(path string, newService SystemdServiceConfig, libDir bool) (SystemdServiceConfig, error) {

	// Expand home directory if path starts with ~
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			path = filepath.Join(homeDir, path[2:])
		}
	}

	// Ensure base directory is absolute
	path, _ = filepath.Abs(path)

	// Build paths
	if libDir {
		path = filepath.Join("/lib/systemd/system", path)
	} else {
		path = filepath.Join("/etc/systemd/system", path)
	}

	newServiceString,err := newService.ToString()

	if err != nil {
		return newService, fmt.Errorf("failed to write service file: %w", err)
	}
	// Write the new service file
	err = SaveFile(path, newServiceString)
	if err != nil {
		return newService, fmt.Errorf("failed to write service file: %w", err)
	}

	return newService, nil
}
