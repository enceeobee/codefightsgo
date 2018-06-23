package regularhell

import (
	"regexp"
	"strings"
)

func programTranslation(solution string, args []string) string {
	argumentVariants := strings.Join(args, "|")
	re := regexp.MustCompile(`([^$a-zA-Z_])(` + argumentVariants + `)(\W)`)
	repl := "$1$$$2$3"
	subSolution := re.ReplaceAllString(solution, repl)
	return re.ReplaceAllString(subSolution, repl)
}
