package main

// CamelToSnakeCase converts a camelCase string to snake_case without additional imports
func CamelToSnakeCase(s string) string {
	// Handle empty string case
	if len(s) == 0 {
		return ""
	}

	// Check if the string contains non-alphabet characters or ends with an uppercase letter
	for i := 0; i < len(s); i++ {
		if s[i] < 'A' || (s[i] > 'Z' && s[i] < 'a') || s[i] > 'z' {
			return s
		}
	}
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return s
	}

	// Convert camelCase to snake_case
	var result []byte
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, s[i])
	}

	// Convert the result to lowercase manually
	for i := 0; i < len(result); i++ {
		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] += 'a' - 'A'
		}
	}

	return string(result)
}

func main() {
	// Sample test cases
	println(CamelToSnakeCase("HelloWorld"))      // hello_world
	println(CamelToSnakeCase("helloWorld"))      // hello_world
	println(CamelToSnakeCase("camelCase"))       // camel_case
	println(CamelToSnakeCase("CAMELtoSnackCASE")) // CAMELtoSnackCASE (invalid camelCase)
	println(CamelToSnakeCase("camelToSnakeCase")) // camel_to_snake_case
	println(CamelToSnakeCase("hey2"))            // hey2 (invalid camelCase)
}
