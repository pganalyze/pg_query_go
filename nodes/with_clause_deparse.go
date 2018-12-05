// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"strings"
)

func (node WithClause) Deparse(ctx DeparseContext) (string, error) {
	out := []string{"WITH"}
	if node.Recursive {
		out = append(out, "RECURSIVE")
	}

	if node.Ctes.Items == nil || len(node.Ctes.Items) == 0 {
		return "", errors.New("cannot have with clause without ctes")
	}

	if cteQueries, err := node.Ctes.DeparseList(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, strings.Join(cteQueries, ", "))
		return strings.Join(out, " "), nil
	}
}
