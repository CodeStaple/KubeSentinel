package llm

import "fmt"

// LLMConnector defines the interface for interacting with a Large Language Model (LLM)
// service. It abstracts the specific LLM provider implementation, allowing different
// models or services to be used by the operator.
type LLMConnector interface {
	// AnalyzeManifest sends a Kubernetes manifest to the LLM for analysis based on
	// predefined security rules, best practices, or specific queries.
	// It returns a string containing the LLM's analysis and any encountered error.
	AnalyzeManifest(manifest string) (string, error)

	// SummarizeAnomalies takes a list of security event descriptions or anomaly data
	// and requests the LLM to provide a concise summary or identify patterns.
	// It returns a string summary and any encountered error.
	SummarizeAnomalies(events []string) (string, error)
}

// DummyLLMConnector is a placeholder implementation of LLMConnector.
// It's used for testing or when a real LLM integration is not available.
// Its methods return predefined, hardcoded responses and do not make actual LLM calls.
type DummyLLMConnector struct{}

// NewDummyLLMConnector creates a new DummyLLMConnector.
func NewDummyLLMConnector() *DummyLLMConnector {
	return &DummyLLMConnector{}
}

// AnalyzeManifest returns a dummy analysis for a given manifest.
// It simulates an LLM analyzing a Kubernetes manifest.
func (c *DummyLLMConnector) AnalyzeManifest(manifest string) (string, error) {
	response := fmt.Sprintf("Dummy LLM analysis of manifest content: %s. Conclusion: The manifest appears to be structured correctly according to dummy checks.", manifest)
	fmt.Printf("LLMConnector (Dummy): AnalyzeManifest called for manifest starting with: %.50s...\n", manifest)
	return response, nil
}

// SummarizeAnomalies returns a dummy summary for given events.
// It simulates an LLM summarizing a list of security anomalies.
func (c *DummyLLMConnector) SummarizeAnomalies(events []string) (string, error) {
	response := fmt.Sprintf("Dummy LLM summary of %d security anomalies. General advice: Investigate all anomalies thoroughly.", len(events))
	fmt.Printf("LLMConnector (Dummy): SummarizeAnomalies called with %d events.\n", len(events))
	return response, nil
}
