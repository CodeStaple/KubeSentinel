package trivy

import (
	"fmt"
	// Import the scanner interface - this should be uncommented when fully integrated
	// "github.com/example/my-operator/pkg/scanners"
)

// DummyTrivyScanner is a placeholder implementation of the scanners.VulnerabilityScanner interface.
// It simulates the behavior of Trivy for vulnerability scanning but returns hardcoded results
// without actually invoking Trivy. This is useful for testing or when Trivy is not available.
type DummyTrivyScanner struct{}

// NewDummyTrivyScanner creates a new DummyTrivyScanner.
func NewDummyTrivyScanner() *DummyTrivyScanner {
	return &DummyTrivyScanner{}
}

// ScanImage returns a dummy vulnerability report for a given image.
// It's meant to implement the scanners.VulnerabilityScanner interface.
// The actual Trivy scanner would perform a real scan of the container image.
func (s *DummyTrivyScanner) ScanImage(imageName string) (string, error) {
	report := fmt.Sprintf("Dummy Trivy scan report for image: %s\n"+
		"-------------------------------------\n"+
		"CVE-2024-0001: HIGH - Outdated library 'libexample' (version 1.0.0, fixed in 1.0.1)\n"+
		"CVE-2024-0002: MEDIUM - Insecure default configuration in '/etc/dummy.conf'\n"+
		"CVE-2024-0003: LOW - Debug endpoint exposed on port 9999\n"+
		"-------------------------------------\n"+
		"Summary: 1 Critical, 0 High, 1 Medium, 1 Low vulnerabilities found by dummy scanner.", imageName)
	fmt.Printf("TrivyScanner (Dummy): ScanImage called for image: %s\n", imageName)
	return report, nil
}
