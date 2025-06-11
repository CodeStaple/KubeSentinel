package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AIPodSecurityAssessmentSpec defines the desired state of AIPodSecurityAssessment.
// It specifies the Pod to assess and the security profile to assess against.
type AIPodSecurityAssessmentSpec struct {
	// PodName is the name of the Pod to assess. This field is required.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	PodName string `json:"podName"`

	// Namespace is the namespace of the Pod to assess. This field is required.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Namespace string `json:"namespace"`

	// SecurityProfile specifies the security profile or policy to assess the Pod against
	// (e.g., "baseline", "restricted" from Pod Security Standards, or a custom profile name).
	// If not specified, a default baseline profile may be used.
	// +optional
	SecurityProfile string `json:"securityProfile,omitempty"`

	// IncludeRecommendations indicates whether to include AI-generated recommendations for improving the Pod's security posture.
	// +optional
	IncludeRecommendations bool `json:"includeRecommendations,omitempty"`
}

// AIPodSecurityAssessmentStatus defines the observed state of AIPodSecurityAssessment.
// It holds the results and status of the Pod security assessment.
type AIPodSecurityAssessmentStatus struct {
	// AssessmentResult contains a summary of the assessment findings (e.g., "Compliant", "NonCompliant", "NeedsReview").
	// +optional
	AssessmentResult string `json:"assessmentResult,omitempty"`

	// Violations lists any security policy violations or deviations from the specified security profile.
	// Each item could be a string describing the violation or a more structured object.
	// +optional
	Violations []string `json:"violations,omitempty"` // Consider a structured type for ViolationDetail later

	// Recommendations provides actionable advice for improving the Pod's security settings.
	// +optional
	Recommendations []string `json:"recommendations,omitempty"`

	// LastAssessmentTime records the timestamp when the assessment was last performed.
	// +optional
	LastAssessmentTime *metav1.Time `json:"lastAssessmentTime,omitempty"`

	// Error stores any error messages encountered during the assessment process.
	// +optional
	Error string `json:"error,omitempty"`

	// ObservedGeneration is the most recent generation observed by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Conditions store the status conditions of the AIPodSecurityAssessment instances
	// +operator-sdk:csv:customresourcedefinitions:type=status
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

// AIPodSecurityAssessment is the Schema for the aipodsecurityassessments API.
// It represents a request to assess a specific Pod against Kubernetes Pod Security Standards (PSS),
// custom security policies, or other security best practices. AI can be used to provide
// context-aware explanations of findings or suggest remediations.
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="PodName",type="string",JSONPath=".spec.podName",description="Name of the Pod assessed"
//+kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".spec.namespace",description="Namespace of the Pod"
//+kubebuilder:printcolumn:name="Profile",type="string",JSONPath=".spec.securityProfile",description="Security profile used"
//+kubebuilder:printcolumn:name="Result",type="string",JSONPath=".status.assessmentResult",description="Overall assessment result"
//+kubebuilder:printcolumn:name="Violations",type="integer",JSONPath=".status.violations.#",description="Number of violations found"
//+kubebuilder:printcolumn:name="LastAssessment",type="date",JSONPath=".status.lastAssessmentTime",description="Time of last assessment"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:resource:path=aipodsecurityassessments,scope=Namespaced,shortName=apsa
type AIPodSecurityAssessment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AIPodSecurityAssessmentSpec   `json:"spec,omitempty"`
	Status AIPodSecurityAssessmentStatus `json:"status,omitempty"`
}

// AIPodSecurityAssessmentList contains a list of AIPodSecurityAssessment
//+kubebuilder:object:root=true
type AIPodSecurityAssessmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPodSecurityAssessment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPodSecurityAssessment{}, &AIPodSecurityAssessmentList{})
}
