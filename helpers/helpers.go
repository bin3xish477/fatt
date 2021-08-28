package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/binexisHATT/fatt/patterns"
	"github.com/dlclark/regexp2"
	"github.com/olekukonko/tablewriter"
)

// FindAllString returns a string slice containing all the
// strings that matched the specified pattern
func FindAllString(re *regexp2.Regexp, s string) []string {
	var matches []string
	m, _ := re.FindStringMatch(s)
	for m != nil {
		matches = append(matches, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return matches
}

// ListPatterns lists all of fatt's available search patterns
func ListPatterns() {
	data := [][]string{}
	for patternName, _ := range patterns.Patterns {
		data = append(data, []string{patternName})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"pattern name"})
	table.SetFooter([]string{fmt.Sprintf("# of Available Patterns: %d", len(patterns.Patterns))})
	//table.SetBorder(false)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

// StringInSlice returns true if `str` is in string slice `slice`
func StringInSlice(str string, slice []string) bool {
	// convert string slice to lowercase
	var temp []string
	for _, s := range slice {
		temp = append(temp, strings.ToLower(s))
	}

	for _, s := range temp {
		if s == str {
			return true
		}
	}
	return false
}

// FileExists returns true is `filename` exists and is not a directory
// returns false otherwise
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
