//go:build windows
// +build windows

package collector

import (
	"os/exec"
	"strings"

	"github.com/shaswatprakash/windows-agent/internal/models"
	"golang.org/x/sys/windows/registry"
)

func RunCISChecks() []models.CISCheck {
	checks := []models.CISCheck{}

	// 1. Firewall profiles enabled
	firewall := runPowerShell(`(Get-NetFirewallProfile | Where {$_.Enabled -eq "True"}).Count`)
	checks = append(checks, models.CISCheck{
		Name:     "Firewall profiles enabled",
		Passed:   strings.TrimSpace(firewall) == "3",
		Evidence: "Profiles enabled: " + firewall,
	})

	// 2. BitLocker enabled
	bitlocker := runPowerShell(`(Get-BitLockerVolume -MountPoint "C:").VolumeStatus`)
	checks = append(checks, models.CISCheck{
		Name:     "BitLocker enabled",
		Passed:   strings.Contains(bitlocker, "FullyEncrypted"),
		Evidence: bitlocker,
	})

	// 3. SMBv1 disabled
	smbv1, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\LanmanServer\Parameters`, registry.READ)
	val, _, _ := smbv1.GetIntegerValue("SMB1")
	checks = append(checks, models.CISCheck{
		Name:     "SMBv1 disabled",
		Passed:   val == 0,
		Evidence: "SMB1=" + string(rune(val)),
	})
	smbv1.Close()

	// 4. RDP NLA enabled
	rdp := runPowerShell(`(Get-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Control\Terminal Server\WinStations\RDP-Tcp").UserAuthentication`)
	checks = append(checks, models.CISCheck{
		Name:     "RDP NLA enabled",
		Passed:   strings.TrimSpace(rdp) == "1",
		Evidence: "UserAuthentication=" + rdp,
	})

	// 5. Password complexity enabled
	pass := runPowerShell(`(Get-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Control\Lsa").PasswordComplexity`)
	checks = append(checks, models.CISCheck{
		Name:     "Password complexity enabled",
		Passed:   strings.TrimSpace(pass) == "1",
		Evidence: "PasswordComplexity=" + pass,
	})

	// 6. Account lockout policy set
	lockout := runPowerShell(`(net accounts | find "lockout threshold")`)
	checks = append(checks, models.CISCheck{
		Name:     "Account lockout policy set",
		Passed:   strings.Contains(lockout, "5"),
		Evidence: lockout,
	})

	// 7. UAC enabled
	uac := runPowerShell(`(Get-ItemProperty -Path "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System").EnableLUA`)
	checks = append(checks, models.CISCheck{
		Name:     "UAC enabled",
		Passed:   strings.TrimSpace(uac) == "1",
		Evidence: "EnableLUA=" + uac,
	})

	// 8. LSA protection enabled
	lsa := runPowerShell(`(Get-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Control\Lsa").RunAsPPL`)
	checks = append(checks, models.CISCheck{
		Name:     "LSA protection enabled",
		Passed:   strings.TrimSpace(lsa) == "1",
		Evidence: "RunAsPPL=" + lsa,
	})

	// 9. Windows Defender enabled
	defender := runPowerShell(`(Get-MpComputerStatus).RealTimeProtectionEnabled`)
	checks = append(checks, models.CISCheck{
		Name:     "Windows Defender enabled",
		Passed:   strings.TrimSpace(defender) == "True",
		Evidence: "RealTimeProtection=" + defender,
	})

	// 10. Automatic updates enabled
	updates := runPowerShell(`(Get-ItemProperty -Path "HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU").NoAutoUpdate`)
	checks = append(checks, models.CISCheck{
		Name:     "Automatic updates enabled",
		Passed:   strings.TrimSpace(updates) == "0",
		Evidence: "NoAutoUpdate=" + updates,
	})

	return checks
}

func runPowerShell(cmd string) string {
	out, err := exec.Command("powershell", "-Command", cmd).Output()
	if err != nil {
		return "Error"
	}
	return string(out)
}
