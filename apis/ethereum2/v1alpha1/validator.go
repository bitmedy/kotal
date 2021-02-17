package v1alpha1

import (
	"github.com/kotalco/kotal/apis/shared"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ValidatorSpec defines the desired state of Validator
type ValidatorSpec struct {
	// Network is the network this validator is validating blocks for
	Network string `json:"network"`
	// Client is the Ethereum 2.0 client to use
	Client Ethereum2Client `json:"client,omitempty"`
	// BeaconEndpoint is the beacon node endpoint
	BeaconEndpoint string `json:"beaconEndpoint"`
	// Graffiti is the text to include in proposed blocks
	Graffiti string `json:"graffiti,omitempty"`
	// Resources is node compute and storage resources
	shared.Resources `json:"resources,omitempty"`
}

// ValidatorStatus defines the observed state of Validator
type ValidatorStatus struct{}

// +kubebuilder:object:root=true

// Validator is the Schema for the validators API
type Validator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ValidatorSpec   `json:"spec,omitempty"`
	Status ValidatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ValidatorList contains a list of Validator
type ValidatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Validator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Validator{}, &ValidatorList{})
}
