package main

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := []string{}
	couples := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	for _, char := range s {
		c := string(char)
		_, ok := couples[c]
		if ok {
			stack = append(stack, c)
			continue
		} else if len(stack) == 0 {
			return false
		} else {
			right, _ := couples[stack[len(stack)-1]]
			if right == c {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}
	if len(stack) == 0 {
		return true
	}
	return false
}
