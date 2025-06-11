package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AINetworkPolicyAuditSpec defines the desired state of AINetworkPolicyAudit.
// It specifies the scope and parameters for auditing Kubernetes NetworkPolicies.
type AINetworkPolicyAuditSpec struct {
	// Namespace to audit network policies in.
	// If empty, the audit may apply to all namespaces the operator has access to, or a default set.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// IncludeBaselinePolicies indicates whether to include checks against a predefined set of baseline network security policies.
	// These baselines could enforce common security practices (e.g., default deny, restricted egress).
	// +optional
	IncludeBaselinePolicies bool `json:"includeBaselinePolicies,omitempty"`

	// CustomAuditRules contains specific rules or checks to perform during the audit, potentially as a structured query or policy definition.
	// This could be a JSON string or a reference to a ConfigMap containing audit rules.
	// +optional
	CustomAuditRules string `json:"customAuditRules,omitempty"` // Consider a more structured type later
}

// AINetworkPolicyAuditStatus defines the observed state of AINetworkPolicyAudit.
// It holds the results and status of the network policy audit.
type AINetworkPolicyAuditStatus struct {
	// AuditResult contains a summary of the audit findings (e.g., "Compliant", "NonCompliant", "PartiallyCompliant").
	// +optional
	AuditResult string `json:"auditResult,omitempty"`

	// NonCompliantPolicies lists any network policies that were found to be non-compliant or problematic.
	// Each item could be a string reference (namespace/name) or a more detailed finding object.
	// +optional
	NonCompliantPolicies []string `json:"nonCompliantPolicies,omitempty"` // Consider a structured type for PolicyFinding later

	// Suggestions provides actionable recommendations for improving network policy posture.
	// +optional
	Suggestions []string `json:"suggestions,omitempty"`

	// LastAuditTime records the timestamp when the audit was last performed.
	// +optional
	LastAuditTime *metav1.Time `json:"lastAuditTime,omitempty"`

	// Error stores any error messages encountered during the audit process.
	// +optional
	Error string `json:"error,omitempty"`

	// ObservedGeneration is the most recent generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions store the status conditions of the AINetworkPolicyAudit instances
	// +operator-sdk:csv:customresourcedefinitions:type=status
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// AINetworkPolicyAudit is the Schema for the ainetworkpolicyaudits API.
// It represents a request to audit Kubernetes NetworkPolicies within the cluster or a specific namespace.
// The audit can check for compliance with best practices, identify overly permissive policies,
// or detect potential security gaps, potentially using AI to interpret policy effectiveness.
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".spec.namespace",description="Namespace audited (if specified)"
//+kubebuilder:printcolumn:name="Result",type="string",JSONPath=".status.auditResult",description="Overall audit result"
//+kubebuilder:printcolumn:name="NonCompliant",type="integer",JSONPath=".status.nonCompliantPolicies.#",description="Number of non-compliant policies"
//+kubebuilder:printcolumn:name="LastAuditTime",type="date",JSONPath=".status.lastAuditTime",description="Time of the last audit"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:resource:path=ainetworkpolicyaudits,scope=Namespaced,shortName=anpa // Can be cluster-scoped if namespace is empty and permissions allow
type AINetworkPolicyAudit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AINetworkPolicyAuditSpec   `json:"spec,omitempty"`
	Status AINetworkPolicyAuditStatus `json:"status,omitempty"`
}

// AINetworkPolicyAuditList contains a list of AINetworkPolicyAudit
//+kubebuilder:object:root=true
type AINetworkPolicyAuditList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AINetworkPolicyAudit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AINetworkPolicyAudit{}, &AINetworkPolicyAuditList{})
}
