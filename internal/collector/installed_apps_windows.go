//go:build windows
// +build windows

package collector

import (
	"github.com/shaswatprakash/windows-agent/internal/models"
	"golang.org/x/sys/windows/registry"
)

// Windows-specific registry logic
func GetInstalledApps() ([]models.Application, error) {
	apps := []models.Application{}
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		`SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Uninstall`,
		registry.READ)
	if err != nil {
		return apps, err
	}
	defer k.Close()

	names, _ := k.ReadSubKeyNames(-1)
	for _, name := range names {
		sub, err := registry.OpenKey(k, name, registry.READ)
		if err != nil {
			continue
		}
		displayName, _, _ := sub.GetStringValue("DisplayName")
		displayVersion, _, _ := sub.GetStringValue("DisplayVersion")
		if displayName != "" {
			apps = append(apps, models.Application{
				Name:    displayName,
				Version: displayVersion,
			})
		}
		sub.Close()
	}
	return apps, nil
}
