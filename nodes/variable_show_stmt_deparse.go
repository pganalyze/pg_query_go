// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
)

func (node VariableShowStmt) Deparse(ctx DeparseContext) (string, error) {
	return fmt.Sprintf("SHOW %s", *node.Name), nil
}
