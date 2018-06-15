package regularhell

import (
	"regexp"
	"strings"
)

func programTranslation(solution string, args []string) string {
	argumentVariants := strings.Join(args, "|")

	// fmt.Println("argumentVariants", argumentVariants)

	// return ""

	re := regexp.MustCompile("([^\\$a-zA-Z])(" + argumentVariants + ")(\\W)")
	repl := "$1$$$2$3"
	subSolution := solution
	return re.ReplaceAllString(subSolution, repl)
}
