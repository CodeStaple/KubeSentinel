package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AIManifestAnalysisSpec defines the desired state of AIManifestAnalysis.
// It specifies the manifest to be analyzed and any configuration for the analysis.
type AIManifestAnalysisSpec struct {
	// Manifest contains the raw Kubernetes manifest content (e.g., YAML or JSON string) to be analyzed.
	// This field is required.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Manifest string `json:"manifest"`

	// LLMProvider specifies which Large Language Model provider to use for the analysis.
	// Examples: "OpenAI", "Groq", "Gemini", "LocalLLM".
	// If not specified, a default provider may be used.
	// +optional
	LLMProvider string `json:"llmProvider,omitempty"`

	// AnalysisType indicates the type or focus of the analysis (e.g., "security", "best-practices", "cost-optimization").
	// This can guide the LLM in its assessment.
	// +optional
	AnalysisType string `json:"analysisType,omitempty"`
}

// AIManifestAnalysisStatus defines the observed state of AIManifestAnalysis.
// It holds the results and status of the manifest analysis.
type AIManifestAnalysisStatus struct {
	// AnalysisResult contains the findings or summary from the LLM analysis.
	// This could be a textual summary, structured data (JSON string), or a reference to a more detailed report.
	// +optional
	AnalysisResult string `json:"analysisResult,omitempty"`

	// Suggestions provides actionable suggestions based on the analysis.
	// Each suggestion might include a description, severity, and recommended action.
	// +optional
	Suggestions []string `json:"suggestions,omitempty"` // Consider a more structured type for suggestions later

	// Error stores any error messages encountered during the analysis process.
	// If this field is populated, it indicates that the analysis may not have completed successfully.
	// +optional
	Error string `json:"error,omitempty"`

	// LastAnalyzedAt records the timestamp when the analysis was last performed.
	// +optional
	LastAnalyzedAt *metav1.Time `json:"lastAnalyzedAt,omitempty"`

	// ObservedGeneration is the most recent generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions store the status conditions of the AIManifestAnalysis instances
	// +operator-sdk:csv:customresourcedefinitions:type=status
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// AIManifestAnalysis is the Schema for the aimanifestanalyses API.
// It represents a request to analyze a Kubernetes manifest using an AI model
// to identify potential issues, suggest improvements, or check for security vulnerabilities.
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Manifest",type="string",JSONPath=".spec.manifest",description="The Kubernetes manifest to be analyzed (summary or reference)"
//+kubebuilder:printcolumn:name="Result",type="string",JSONPath=".status.analysisResult",description="Summary of the analysis result"
//+kubebuilder:printcolumn:name="Error",type="string",JSONPath=".status.error",description="Any errors during analysis"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:resource:path=aimanifestanalyses,scope=Namespaced,shortName=ama
type AIManifestAnalysis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AIManifestAnalysisSpec   `json:"spec,omitempty"`
	Status AIManifestAnalysisStatus `json:"status,omitempty"`
}

// AIManifestAnalysisList contains a list of AIManifestAnalysis
//+kubebuilder:object:root=true
type AIManifestAnalysisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIManifestAnalysis `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIManifestAnalysis{}, &AIManifestAnalysisList{})
}
