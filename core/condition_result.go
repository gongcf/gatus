package core

// ConditionResult result of a Condition
type ConditionResult struct {
	// Condition that was evaluated
	Condition string `json:"condition"`

	Result interface{} `json:"result,omitempty"`
	// Success whether the condition was met (successful) or not (failed)
	Success bool `json:"success"`
}
