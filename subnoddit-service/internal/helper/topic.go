package helper

import "strings"

func SplitTopicIDs(input string) []string {
	// Split by comma
	topicIDs := strings.Split(input, ",")

	// Optional: Trim whitespace from each ID
	for i, id := range topicIDs {
		topicIDs[i] = strings.TrimSpace(id)
	}

	return topicIDs
}
