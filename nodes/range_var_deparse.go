// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node RangeVar) Deparse(ctx DeparseContext) (string, error) {
	out := make([]string, 0)
	if !node.Inh {
		out = append(out, "ONLY")
	}

	if node.Schemaname != nil && len(*node.Schemaname) > 0 {
		out = append(out, fmt.Sprintf(`"%s"."%s"`, *node.Schemaname, *node.Relname))
	} else {
		out = append(out, fmt.Sprintf(`"%s"`, *node.Relname))
	}

	if node.Alias != nil {
		if str, err := node.Alias.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	return strings.Join(out, " "), nil
}
