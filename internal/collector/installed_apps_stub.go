//go:build !windows
// +build !windows

package collector

import "github.com/shaswatprakash/windows-agent/internal/models"

// Mac/Linux placeholder for installed apps
func GetInstalledApps() ([]models.Application, error) {
	apps := []models.Application{
		{Name: "MockApp", Version: "1.0.0"},
		{Name: "Chrome", Version: "120.0.0"},
	}
	return apps, nil
}
