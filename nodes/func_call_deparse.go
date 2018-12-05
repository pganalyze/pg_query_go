// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node FuncCall) Deparse(ctx DeparseContext) (string, error) {
	out := make([]string, 0)

	args := make([]string, len(node.Args.Items))
	args, err := node.Args.DeparseList(DeparseContextNone)
	if err != nil {
		return "", err
	}

	if node.AggStar {
		args = append(args, "*")
	}

	difference := func(slice1 []string, slice2 []string) []string {
		var diff []string
		// Loop two times, first to find slice1 strings not in slice2,
		// second loop to find slice2 strings not in slice1
		for i := 0; i < 2; i++ {
			for _, s1 := range slice1 {
				found := false
				for _, s2 := range slice2 {
					if s1 == s2 {
						found = true
						break
					}
				}
				// String not found. We add it to return slice
				if !found {
					diff = append(diff, s1)
				}
			}
			// Swap the slices, only if it was the first loop
			if i == 0 {
				slice1, slice2 = slice2, slice1
			}
		}
		return diff
	}

	name := ""
	if names, err := node.Funcname.DeparseList(DeparseContextNone); err != nil {
		return "", err
	} else {
		name = strings.Join(difference([]string{"pg_catalog"}, names), ".")
	}

	distinct := ""
	if node.AggDistinct {
		distinct = "DISTINCT "
	}

	out = append(out, fmt.Sprintf("%s(%s%s)", name, distinct, strings.Join(args, ", ")))

	if node.Over != nil {
		if over, err := node.Over.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, fmt.Sprintf("OVER (%s)", over))
		}
	}

	return strings.Join(out, " "), nil
}
