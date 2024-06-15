package domain

// DetectionRuleType is a custom type to represent the type of detection rule.
type DetectionRuleType string

// Constants for DetectionRuleType to ensure type safety and clarity.
const (
	FileRule      DetectionRuleType = "file"      // Detects the app based on a file's presence or properties
	RegistryRule  DetectionRuleType = "registry"  // Detects the app based on registry entries
	ScriptRule    DetectionRuleType = "script"    // Detects the app based on the output of a custom script
	// Add more rule types as needed
)

// DetectionRule defines a rule to detect if the app is already installed.
type DetectionRule struct {
	RuleType DetectionRuleType // Type of detection rule, using the custom type
	Path     string            // Path to check for the rule (file path, registry path, etc.)
	Value    string            // Value to check against in the path (file version, registry value, script output, etc.)
}

