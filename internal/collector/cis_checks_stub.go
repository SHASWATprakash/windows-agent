//go:build !windows
// +build !windows

package collector

import "github.com/shaswatprakash/windows-agent/internal/models"

func RunCISChecks() []models.CISCheck {
	return []models.CISCheck{
		{Name: "Firewall profiles enabled", Passed: true, Evidence: "Mocked: all profiles enabled"},
		{Name: "BitLocker enabled", Passed: false, Evidence: "Mocked: system drive not encrypted"},
		{Name: "SMBv1 disabled", Passed: true, Evidence: "Mocked: SMBv1 service disabled"},
		{Name: "RDP NLA enabled", Passed: true, Evidence: "Mocked: NLA required for RDP"},
		{Name: "Password complexity enabled", Passed: true, Evidence: "Mocked: complexity enforced"},
		{Name: "Account lockout policy set", Passed: true, Evidence: "Mocked: lockout threshold = 5"},
		{Name: "UAC enabled", Passed: true, Evidence: "Mocked: UAC always notify"},
		{Name: "LSA protection enabled", Passed: false, Evidence: "Mocked: LSA not protected"},
		{Name: "Windows Defender enabled", Passed: true, Evidence: "Mocked: real-time protection active"},
		{Name: "Automatic updates enabled", Passed: true, Evidence: "Mocked: updates auto-download"},
	}
}
