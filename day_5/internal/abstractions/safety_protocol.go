package abstractions

type SafetyProtocol struct {
	rules []OrderingRule
}

func NewSafetyProtocol(
	rules []OrderingRule,
) *SafetyProtocol {
	return &SafetyProtocol{rules}
}
