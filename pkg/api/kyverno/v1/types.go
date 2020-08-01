package v1

import (
	authenticationv1 "k8s.io/api/authentication/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//GenerateRequest is a request to process generate rule
type GenerateRequest struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	// Spec is the information to identify the generate request
	Spec GenerateRequestSpec `json:"spec" yaml:"spec"`
	// Status contains statistics related to generate request
	Status GenerateRequestStatus `json:"status" yaml:"status"`
}

//GenerateRequestSpec stores the request specification
type GenerateRequestSpec struct {
	// Specifies the name of the policy
	Policy string `json:"policy" yaml:"policy"`
	// ResourceSpec is the information to identify the generate request
	Resource ResourceSpec `json:"resource" yaml:"resource"`
	// Context ...
	Context GenerateRequestContext `json:"context"`
}

//GenerateRequestContext stores the context to be shared
type GenerateRequestContext struct {
	UserRequestInfo RequestInfo `json:"userInfo,omitempty" yaml:"userInfo,omitempty"`
}

// RequestInfo contains permission info carried in an admission request
type RequestInfo struct {
	// Roles is a list of possible role send the request
	Roles []string `json:"roles" yaml:"roles"`
	// ClusterRoles is a list of possible clusterRoles send the request
	ClusterRoles []string `json:"clusterRoles" yaml:"clusterRoles"`
	// UserInfo is the userInfo carried in the admission request
	AdmissionUserInfo authenticationv1.UserInfo `json:"userInfo" yaml:"userInfo"`
}

//GenerateRequestStatus stores the status of generated request
type GenerateRequestStatus struct {
	// State represents state of the generate request
	State GenerateRequestState `json:"state" yaml:"state"`
	// Specifies request status message
	// +optional
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	// This will track the resources that are generated by the generate Policy
	// Will be used during clean up resources
	GeneratedResources []ResourceSpec `json:"generatedResources,omitempty" yaml:"generatedResources,omitempty"`
}

//GenerateRequestState defines the state of
type GenerateRequestState string

const (
	//Pending - the Request is yet to be processed or resource has not been created
	Pending GenerateRequestState = "Pending"
	//Failed - the Generate Request Controller failed to process the rules
	Failed GenerateRequestState = "Failed"
	//Completed - the Generate Request Controller created resources defined in the policy
	Completed GenerateRequestState = "Completed"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//GenerateRequestList stores the list of generate requests
type GenerateRequestList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []GenerateRequest `json:"items" yaml:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPolicy ...
type ClusterPolicy Policy

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPolicyList ...
type ClusterPolicyList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []ClusterPolicy `json:"items" yaml:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPolicyViolation represents cluster-wide violations
type ClusterPolicyViolation PolicyViolationTemplate

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPolicyViolationList ...
type ClusterPolicyViolationList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []ClusterPolicyViolation `json:"items" yaml:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PolicyViolation represents namespaced violations
type PolicyViolation PolicyViolationTemplate

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PolicyViolationList ...
type PolicyViolationList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata" yaml:"metadata"`
	Items           []PolicyViolation `json:"items" yaml:"items"`
}

// Policy contains rules to be applied to created resources
type Policy struct {
	metav1.TypeMeta   `json:",inline,omitempty" yaml:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	// Spec is the information to identify the policy
	Spec Spec `json:"spec" yaml:"spec"`
	// Status contains statistics related to policy
	Status PolicyStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// Spec describes policy behavior by its rules
type Spec struct {
	// Rules contains the list of rules to be applied to resources
	Rules []Rule `json:"rules,omitempty" yaml:"rules,omitempty"`
	// ValidationFailureAction provides choice to enforce rules to resources during policy violations.
	// Default value is "audit".
	ValidationFailureAction string `json:"validationFailureAction,omitempty" yaml:"validationFailureAction,omitempty"`
	// Background provides choice for applying rules to existing resources.
	// Default value is "true".
	Background *bool `json:"background,omitempty" yaml:"background,omitempty"`
}

// Rule is set of mutation, validation and generation actions
// for the single resource description
type Rule struct {
	// Specifies rule name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Specifies resources for which the rule has to be applied.
	// If it's defined, "kind" inside MatchResources block is required.
	// +optional
	MatchResources MatchResources `json:"match,omitempty" yaml:"match,omitempty"`
	// Specifies resources for which rule can be excluded
	// +optional
	ExcludeResources ExcludeResources `json:"exclude,omitempty" yaml:"exclude,omitempty"`
	// Allows controlling policy rule execution
	// +optional
	Conditions []Condition `json:"preconditions,omitempty" yaml:"preconditions,omitempty"`
	// Specifies patterns to mutate resources
	// +optional
	Mutation Mutation `json:"mutate,omitempty" yaml:"mutate,omitempty"`
	// Specifies patterns to validate resources
	// +optional
	Validation Validation `json:"validate,omitempty" yaml:"validate,omitempty"`
	// Specifies patterns to create additional resources
	// +optional
	Generation Generation `json:"generate,omitempty" yaml:"generate,omitempty"`
}

//Condition defines the evaluation condition
type Condition struct {
	// Key contains key to compare
	Key interface{} `json:"key,omitempty" yaml:"key,omitempty"`
	// Operator to compare against value
	Operator ConditionOperator `json:"operator,omitempty" yaml:"operator,omitempty"`
	// Value to be compared
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`
}

// ConditionOperator defines the type for condition operator
type ConditionOperator string

const (
	//Equal for Equal operator
	Equal  ConditionOperator = "Equal"
	Equals ConditionOperator = "Equals"
	//NotEqual for NotEqual operator
	NotEqual  ConditionOperator = "NotEqual"
	NotEquals ConditionOperator = "NotEquals"
	//In for In operator
	In ConditionOperator = "In"
	//NotIn for NotIn operator
	NotIn ConditionOperator = "NotIn"
)

//MatchResources contains resource description of the resources that the rule is to apply on
type MatchResources struct {
	// Specifies user information
	UserInfo `json:",omitempty" yaml:",omitempty"`
	// Specifies resources to which rule is applied
	ResourceDescription `json:"resources,omitempty" yaml:"resources,omitempty"`
}

//ExcludeResources container resource description of the resources that are to be excluded from the applying the policy rule
type ExcludeResources struct {
	// Specifies user information
	UserInfo `json:",omitempty" yaml:",omitempty"`
	// Specifies resources to which rule is excluded
	ResourceDescription `json:"resources,omitempty" yaml:"resources,omitempty"`
}

// UserInfo filter based on users
type UserInfo struct {
	// Specifies list of namespaced role names
	Roles []string `json:"roles,omitempty" yaml:"roles,omitempty"`
	// Specifies list of cluster wide role names
	ClusterRoles []string `json:"clusterRoles,omitempty" yaml:"clusterRoles,omitempty"`
	// Specifies list of subject names like users, user groups, and service accounts
	Subjects []rbacv1.Subject `json:"subjects,omitempty" yaml:"subjects,omitempty"`
}

// ResourceDescription describes the resource to which the PolicyRule will be applied.
type ResourceDescription struct {
	// Specifies list of resource kind
	Kinds []string `json:"kinds,omitempty" yaml:"kinds,omitempty"`
	// Specifies name of the resource
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Specifies list of namespaces
	Namespaces []string `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
	// Specifies the set of selectors
	Selector *metav1.LabelSelector `json:"selector,omitempty" yaml:"selector,omitempty"`
}

// Mutation describes the way how Mutating Webhook will react on resource creation
type Mutation struct {
	// Specifies overlay patterns
	Overlay interface{} `json:"overlay,omitempty" yaml:"overlay,omitempty"`
	// Specifies JSON Patch
	Patches []Patch `json:"patches,omitempty" yaml:"patches,omitempty"`
}

// +k8s:deepcopy-gen=false

// Patch declares patch operation for created object according to RFC 6902
type Patch struct {
	// Specifies path of the resource
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
	// Specifies operations supported by JSON Patch.
	// i.e:- add, replace and delete
	Operation string `json:"op,omitempty" yaml:"op,omitempty"`
	// Specifies the value to be applied
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`
}

// Validation describes the way how Validating Webhook will check the resource on creation
type Validation struct {
	// Specifies message to be displayed on validation policy violation
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	// Specifies validation pattern
	Pattern interface{} `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	// Specifies list of validation patterns
	AnyPattern []interface{} `json:"anyPattern,omitempty" yaml:"anyPattern,omitempty"`
	// Specifies conditions to deny validation
	Deny *Deny `json:"deny,omitempty" yaml:"deny,omitempty"`
}

type Deny struct {
	// Specifies set of condition to deny validation
	Conditions []Condition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}

// Generation describes which resources will be created when other resource is created
type Generation struct {
	ResourceSpec
	// To keep resources synchronized with source resource
	Synchronize bool `json:"synchronize,omitempty" yaml:"synchronize,omitempty"`
	// Data ...
	Data interface{} `json:"data,omitempty" yaml:"data,omitempty"`
	// To clone resource from other resource
	Clone CloneFrom `json:"clone,omitempty" yaml:"clone,omitempty"`
}

// CloneFrom - location of the resource
// which will be used as source when applying 'generate'
type CloneFrom struct {
	// Specifies resource namespace
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	// Specifies name of the resource
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// PolicyStatus mostly contains statistics related to policy
type PolicyStatus struct {
	// average time required to process the policy rules on a resource
	AvgExecutionTime string `json:"averageExecutionTime,omitempty" yaml:"averageExecutionTime,omitempty"`
	// number of violations created by this policy
	ViolationCount int `json:"violationCount,omitempty" yaml:"violationCount,omitempty"`
	// Count of rules that failed
	RulesFailedCount int `json:"rulesFailedCount,omitempty" yaml:"rulesFailedCount,omitempty"`
	// Count of rules that were applied
	RulesAppliedCount int `json:"rulesAppliedCount,omitempty" yaml:"rulesAppliedCount,omitempty"`
	// Count of resources that were blocked for failing a validate, across all rules
	ResourcesBlockedCount int `json:"resourcesBlockedCount,omitempty" yaml:"resourcesBlockedCount,omitempty"`
	// Count of resources that were successfully mutated, across all rules
	ResourcesMutatedCount int `json:"resourcesMutatedCount,omitempty" yaml:"resourcesMutatedCount,omitempty"`
	// Count of resources that were successfully generated, across all rules
	ResourcesGeneratedCount int `json:"resourcesGeneratedCount,omitempty" yaml:"resourcesGeneratedCount,omitempty"`

	Rules []RuleStats `json:"ruleStatus,omitempty" yaml:"ruleStatus,omitempty"`
}

//RuleStats provides status per rule
type RuleStats struct {
	// Rule name
	Name string `json:"ruleName" yaml:"ruleName"`
	// average time require to process the rule
	ExecutionTime string `json:"averageExecutionTime,omitempty" yaml:"averageExecutionTime,omitempty"`
	// number of violations created by this rule
	ViolationCount int `json:"violationCount,omitempty" yaml:"violationCount,omitempty"`
	// Count of rules that failed
	FailedCount int `json:"failedCount,omitempty" yaml:"failedCount,omitempty"`
	// Count of rules that were applied
	AppliedCount int `json:"appliedCount,omitempty" yaml:"appliedCount,omitempty"`
	// Count of resources for whom update/create api requests were blocked as the resource did not satisfy the policy rules
	ResourcesBlockedCount int `json:"resourcesBlockedCount,omitempty" yaml:"resourcesBlockedCount,omitempty"`
	// Count of resources that were successfully mutated
	ResourcesMutatedCount int `json:"resourcesMutatedCount,omitempty" yaml:"resourcesMutatedCount,omitempty"`
	// Count of resources that were successfully generated
	ResourcesGeneratedCount int `json:"resourcesGeneratedCount,omitempty" yaml:"resourcesGeneratedCount,omitempty"`
}

// PolicyList is a list of Policy resources

// PolicyViolationTemplate stores the information regarinding the resources for which a policy failed to apply
type PolicyViolationTemplate struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" `
	Spec              PolicyViolationSpec   `json:"spec" yaml:"spec"`
	Status            PolicyViolationStatus `json:"status" yaml:"status"`
}

// PolicyViolationSpec describes policy behavior by its rules
type PolicyViolationSpec struct {
	// Specifies name of the policy
	Policy       string `json:"policy" yaml:"policy"`
	ResourceSpec `json:"resource" yaml:"resource"`
	// Specifies list of violated rule
	ViolatedRules []ViolatedRule `json:"rules" yaml:"rules"`
}

// ResourceSpec information to identify the resource
type ResourceSpec struct {
	// Specifies resource kind
	// +optional
	Kind string `json:"kind,omitempty" yaml:"kind,omitempty"`
	// Specifies resource namespace
	// +optional
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	// Specifies resource name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// ViolatedRule stores the information regarding the rule
type ViolatedRule struct {
	// Specifies violated rule name
	Name string `json:"name" yaml:"name"`
	// Specifies violated rule type
	Type string `json:"type" yaml:"type"`
	// Specifies violation message
	Message string `json:"message" yaml:"message"`
}

//PolicyViolationStatus provides information regarding policyviolation status
// status:
//		LastUpdateTime : the time the policy violation was updated
type PolicyViolationStatus struct {
	// LastUpdateTime : the time the policy violation was updated
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty" yaml:"lastUpdateTime,omitempty"`
}
