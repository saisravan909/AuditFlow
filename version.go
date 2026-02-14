package main

// GetVersion returns the current alpha version of AuditFlow
func GetVersion() string {
	return "0.1.0-alpha"
}

func main() {
	println("AuditFlow: " + GetVersion())
}
