package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AISecretsAuditSpec defines the desired state of AISecretsAudit.
// It specifies the scope and parameters for auditing Kubernetes Secrets.
type AISecretsAuditSpec struct {
	// Namespace to audit secrets in.
	// If empty, the audit may apply to all namespaces the operator has access to, or a default set.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// IncludeExternalSecrets indicates whether to audit secrets managed by external secret operators (e.g., ExternalSecretsOperator, Vault injector).
	// This might involve checking the configuration or status of such external systems if accessible.
	// +optional
	IncludeExternalSecrets bool `json:"includeExternalSecrets,omitempty"`

	// AuditType specifies the type of audit to perform on secrets (e.g., "WeakPassword", "UnusedSecret", "ExposedInEnvVar").
	// +optional
	AuditType string `json:"auditType,omitempty"`

	// RiskThreshold defines a minimum risk level to report (e.g., "High", "Critical").
	// Secrets with findings below this threshold might be omitted from the report.
	// +optional
	RiskThreshold string `json:"riskThreshold,omitempty"`
}

// AISecretsAuditStatus defines the observed state of AISecretsAudit.
// It holds the results and status of the secrets audit.
type AISecretsAuditStatus struct {
	// AuditResult contains a summary of the audit findings (e.g., "Found 3 risky secrets", "No issues found").
	// +optional
	AuditResult string `json:"auditResult,omitempty"`

	// PotentiallyLeakedSecrets lists secrets that are identified as potentially leaked or insecurely configured.
	// Each item could be a string reference (namespace/name) or a more detailed finding object.
	// +optional
	PotentiallyLeakedSecrets []string `json:"potentiallyLeakedSecrets,omitempty"` // Consider a structured type for SecretFinding later

	// Recommendations provides actionable advice for improving secrets management based on the audit.
	// +optional
	Recommendations []string `json:"recommendations,omitempty"`

	// LastAuditTime records the timestamp when the audit was last performed.
	// +optional
	LastAuditTime *metav1.Time `json:"lastAuditTime,omitempty"`

	// Error stores any error messages encountered during the audit process.
	// +optional
	Error string `json:"error,omitempty"`

	// ObservedGeneration is the most recent generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions store the status conditions of the AISecretsAudit instances
	// +operator-sdk:csv:customresourcedefinitions:type=status
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// AISecretsAudit is the Schema for the aisecretsaudits API.
// It represents a request to audit Kubernetes Secrets for potential security risks,
// such as weak secrets, secrets exposed in environment variables, or unused secrets.
// AI can be used to analyze secret content or usage patterns for anomalies.
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".spec.namespace",description="Namespace audited (if specified)"
//+kubebuilder:printcolumn:name="Result",type="string",JSONPath=".status.auditResult",description="Overall audit result"
//+kubebuilder:printcolumn:name="LeakedSecrets",type="integer",JSONPath=".status.potentiallyLeakedSecrets.#",description="Number of potentially leaked/risky secrets"
//+kubebuilder:printcolumn:name="LastAuditTime",type="date",JSONPath=".status.lastAuditTime",description="Time of the last audit"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:resource:path=aisecretsaudits,scope=Namespaced,shortName=asaud // Can be cluster-scoped if namespace is empty and permissions allow
type AISecretsAudit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AISecretsAuditSpec   `json:"spec,omitempty"`
	Status AISecretsAuditStatus `json:"status,omitempty"`
}

// AISecretsAuditList contains a list of AISecretsAudit
//+kubebuilder:object:root=true
type AISecretsAuditList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AISecretsAudit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AISecretsAudit{}, &AISecretsAuditList{})
}
