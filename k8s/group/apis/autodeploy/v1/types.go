package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AutoDeployUnit struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec AutoDeployUnitSpec `json:"spec"`
}

// AutoDeployUnitSpec todo
type AutoDeployUnitSpec struct {
	Image string `json:"image"`
	// +optional
	Cmd []string `json:"cmd"`
	// +optional
	Url string `json:"url"`
	// +optional
	Env []corev1.EnvVar `json:"env"`
	// +optional
	Sync bool `json:"sync"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AutoDeployUnitList todo
type AutoDeployUnitList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AutoDeployUnit `json:"items"`
}

// AutoDeployUnitStatus todo
// type AutoDeployUnitStatus struct {
// 	UpdateTime metav1.Time       `json:"updatetime,omitempty"`
// 	MetaList   []metav1.TypeMeta `json:"metalist,omitempty"`
// }
