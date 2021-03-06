package goodhosts

import (
	"fmt"
	"strings"
)

func itemInSlice(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}

	return false
}

func sliceContainsItem(item string, list []string) bool {
	for _, i := range list {
		if strings.Contains(i, item) {
			return true
		}
	}

	return false
}

func buildRawLine(ip string, hosts []string) string {
	output := ip
	for _, host := range hosts {
		output = fmt.Sprintf("%s %s", output, host)
	}

	return output
}
