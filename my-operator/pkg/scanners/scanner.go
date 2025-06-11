package scanners

// VulnerabilityScanner defines the interface for a tool or service that scans
// container images for known vulnerabilities. This allows the operator to abstract
// the specific scanning tool being used (e.g., Trivy, Clair, Grype).
type VulnerabilityScanner interface {
	// ScanImage takes an image name (e.g., "nginx:latest" or "myrepo/myimage:v1.2.3")
	// and returns a report of found vulnerabilities as a string (e.g., JSON, plain text summary)
	// and any error encountered during the scan.
	ScanImage(imageName string) (string, error)
}
