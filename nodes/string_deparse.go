// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node String) Deparse(ctx DeparseContext) (string, error) {
	switch ctx {
	case DeparseContextAConst:
		return fmt.Sprintf("'%s'", strings.Replace(node.Str, "'", "''", -1)), nil
	case DeparseContextFuncCall, DeparseContextTypeName, DeparseContextOperator:
		return node.Str, nil
	default:
		return fmt.Sprintf(`"%s"`, strings.Replace(node.Str, `"`, `""`, -1)), nil
	}
}
