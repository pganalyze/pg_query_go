// Auto-generated - DO NOT EDIT

package pg_query

import (
	"strings"
)

func (node VariableSetStmt) Deparse(ctx DeparseContext) (string, error) {
	out := []string{"SET"}
	if node.IsLocal {
		out = append(out, "LOCAL")
	}
	out = append(out, *node.Name)
	out = append(out, "TO")
	if args, err := node.Args.DeparseList(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, args...)
	}
	return strings.Join(out, " "), nil
}
