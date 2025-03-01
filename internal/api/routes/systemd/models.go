package systemd

type CreateSystemDModel struct {
	ServiceName string `json:"serviceName"`
	BaseDir  string `json:"baseDir"`
	ExecCommand string `json:"execCommand"`
	LibDir	bool `json:"libDir"`
}
type CreateSystemDResponseModel struct {
	CreateSystemDModel `json:"createSystemDModel"`
	ServiceFile string `json:"serviceFile"`
	CreatedOrUpdated string `json:"createdOrUpdated"`
}

type ListSystedDServicesResponse struct {

}
