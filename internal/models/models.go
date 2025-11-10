package models

type Application struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CISCheck struct {
	Name     string `json:"name"`
	Passed   bool   `json:"passed"`
	Evidence string `json:"evidence"`
}

type HostData struct {
	Hostname     string        `json:"hostname"`
	Applications []Application `json:"applications"`
	CISChecks    []CISCheck    `json:"cis_checks"`
}
