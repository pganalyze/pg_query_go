// Auto-generated - DO NOT EDIT

package pg_query

import (
	"strings"
)

func (node ColumnDef) Deparse(ctx DeparseContext) (string, error) {
	out := []string{*node.Colname}

	if str, err := node.TypeName.Deparse(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, str)
	}

	if node.RawDefault != nil {
		out = append(out, "USING")
		if str, err := node.RawDefault.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	if node.Constraints.Items != nil && len(node.Constraints.Items) > 0 {
		constraints := make([]string, len(node.Constraints.Items))
		for i, constraint := range node.Constraints.Items {
			if str, err := constraint.Deparse(DeparseContextNone); err != nil {
				return "", err
			} else {
				constraints[i] = str
			}
		}
		out = append(out, constraints...)
	}
	return strings.Join(out, " "), nil
}
