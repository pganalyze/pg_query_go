// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"fmt"
	"strings"
)

func (node ColumnRef) Deparse(ctx DeparseContext) (string, error) {
	if node.Fields.Items == nil || len(node.Fields.Items) == 0 {
		return "", errors.New("columnref cannot have 0 fields")
	}
	out := make([]string, len(node.Fields.Items))
	for i, field := range node.Fields.Items {
		switch f := field.(type) {
		case String:
			out[i] = fmt.Sprintf(`"%s"`, f.Str)
		default:
			if str, err := field.Deparse(DeparseContextNone); err != nil {
				return "", err
			} else {
				out[i] = str
			}
		}
	}
	return strings.Join(out, "."), nil
}
