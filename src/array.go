package main

/* Returns the index of s if array contains s. Returns -1 if s is not contained within array. */
func ArrayIndexOf(array []string, s string) int {
	for index := range array {
		if array[index] == s {
			return index
		}
	}
	return -1
}

/* Makes a pretty string representation of a string array */
func ArrayToString(array []string) string {
	if len(array) == 0 {
		return "[]"
	}

	if len(array) == 1 {
		return "[ " + array[0] + " ]"
	}

	output := "[ "
	lastIndex := len(array) - 1
	for i := 0; i < lastIndex; i++ {
		output += array[i] + ", "

	}
	output += array[lastIndex] + " ]"
	return output
}

/* Returns a map where keys are the strings, and values are the indicies. */
func ArrayToMap(array []string) map[string]int {
	m := make(map[string]int)
	for index, element := range array {
		m[element] = index
	}
	return m
}

/* Removes any elements in blacklist from array. */
func ArrayBlacklist(array []string, blacklist []string) []string {
	blackMap := ArrayToMap(array)
	var output []string
	for _, element := range array {
		if _, contains := blackMap[element]; !contains {
			output = append(output, element)
		}
	}
	return output
}
