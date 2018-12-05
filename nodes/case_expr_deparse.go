// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"strings"
)

func (node CaseExpr) Deparse(ctx DeparseContext) (string, error) {
	out := []string{"CASE"}

	if node.Arg != nil {
		if str, err := node.Arg.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	if node.Args.Items == nil || len(node.Args.Items) == 0 {
		return "", errors.New("case expression cannot have 0 arguments")
	}

	if args, err := node.Args.DeparseList(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, args...)
	}

	if node.Defresult != nil {
		out = append(out, "ELSE")
		if str, err := node.Defresult.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	out = append(out, "END")
	return strings.Join(out, " "), nil
}
