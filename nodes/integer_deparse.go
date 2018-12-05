// Auto-generated - DO NOT EDIT

package pg_query

import (
	"strconv"
)

func (node Integer) Deparse(ctx DeparseContext) (string, error) {
	return strconv.FormatInt(node.Ival, 10), nil
}
