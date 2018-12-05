// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node Alias) Deparse(ctx DeparseContext) (string, error) {
	if node.Colnames.Items != nil && len(node.Colnames.Items) > 0 {
		if colNames, err := node.Colnames.DeparseList(DeparseContextNone); err != nil {
			return "", err
		} else {
			cols := strings.Join(colNames, ", ")
			return fmt.Sprintf(`%s (%s)`, *node.Aliasname, cols), nil
		}
	} else {
		return *node.Aliasname, nil
	}
}
