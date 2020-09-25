/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodeRoleControllerSpec defines the desired state of NodeRoleController
type NodeRoleControllerSpec struct {
	Roles []NodeRole `json:"roles"`
}

// NodeRole ...
type NodeRole struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value string `json:"value"`
}

// NodeRoleControllerStatus defines the observed state of NodeRoleController
type NodeRoleControllerStatus struct {
	NodeLabelStatus string `json:"nodelabelstatus"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// NodeRoleController is the Schema for the noderolecontrollers API
type NodeRoleController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeRoleControllerSpec   `json:"spec,omitempty"`
	Status NodeRoleControllerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeRoleControllerList contains a list of NodeRoleController
type NodeRoleControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeRoleController `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeRoleController{}, &NodeRoleControllerList{})
}
