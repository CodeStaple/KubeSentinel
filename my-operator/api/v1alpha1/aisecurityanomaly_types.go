package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AISecurityAnomalySpec defines the desired state of AISecurityAnomaly.
// It describes the details of a detected security anomaly.
type AISecurityAnomalySpec struct {
	// Description provides a human-readable description of the detected security anomaly.
	// This field is required.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Description string `json:"description"`

	// Severity indicates the severity level of the anomaly (e.g., "Critical", "High", "Medium", "Low", "Informational").
	// This field is required.
	// +kubebuilder:validation:Required
	Severity string `json:"severity"`

	// SourceComponent indicates the component or system that reported the anomaly (e.g., "RuntimeMonitor", "IDS", "LogAnalysis").
	// +optional
	SourceComponent string `json:"sourceComponent,omitempty"`

	// AffectedResource refers to the Kubernetes resource primarily affected by this anomaly, if applicable.
	// Format: namespace/kind/name (e.g., "default/pod/my-malicious-pod").
	// +optional
	AffectedResource string `json:"affectedResource,omitempty"`

	// AnomalyType is a category for the anomaly (e.g., "SuspiciousNetworkActivity", "PrivilegeEscalationAttempt", "UnexpectedProcess").
	// +optional
	AnomalyType string `json:"anomalyType,omitempty"`
}

// AISecurityAnomalyStatus defines the observed state of AISecurityAnomaly.
// It tracks the lifecycle and investigation status of the anomaly.
type AISecurityAnomalyStatus struct {
	// State indicates the current state of the anomaly (e.g., "New", "Investigating", "Mitigating", "Resolved", "FalsePositive").
	// +optional
	State string `json:"state,omitempty"`

	// FirstObservedAt is the timestamp when the anomaly was first detected.
	// +optional
	FirstObservedAt *metav1.Time `json:"firstObservedAt,omitempty"`

	// LastObservedAt is the timestamp when the anomaly was last observed or updated.
	// +optional
	LastObservedAt *metav1.Time `json:"lastObservedAt,omitempty"`

	// AssignedTo indicates the user or team assigned to investigate this anomaly.
	// +optional
	AssignedTo string `json:"assignedTo,omitempty"`

	// ResolutionDetails provides information about how the anomaly was resolved, if applicable.
	// +optional
	ResolutionDetails string `json:"resolutionDetails,omitempty"`

	// ObservedGeneration is the most recent generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions store the status conditions of the AISecurityAnomaly instances
	// +operator-sdk:csv:customresourcedefinitions:type=status
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// AISecurityAnomaly is the Schema for the aisecurityanomalies API.
// It represents a detected security anomaly within the Kubernetes cluster, potentially identified
// by AI-driven analysis of runtime behavior, logs, or other security signals.
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Description",type="string",JSONPath=".spec.description",description="Description of the anomaly"
//+kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".spec.severity",description="Severity of the anomaly"
//+kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state",description="Current state of the anomaly"
//+kubebuilder:printcolumn:name="AffectedResource",type="string",JSONPath=".spec.affectedResource",description="Affected K8s resource"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:resource:path=aisecurityanomalies,scope=Namespaced,shortName=asa
type AISecurityAnomaly struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AISecurityAnomalySpec   `json:"spec,omitempty"`
	Status AISecurityAnomalyStatus `json:"status,omitempty"`
}

// AISecurityAnomalyList contains a list of AISecurityAnomaly
//+kubebuilder:object:root=true
type AISecurityAnomalyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AISecurityAnomaly `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AISecurityAnomaly{}, &AISecurityAnomalyList{})
}
