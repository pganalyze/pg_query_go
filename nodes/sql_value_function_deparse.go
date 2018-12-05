// Auto-generated - DO NOT EDIT

package pg_query

func (node SQLValueFunction) Deparse(ctx DeparseContext) (string, error) {
	switch node.Op {
	case SVFOP_CURRENT_TIMESTAMP:
		return "CURRENT_TIMESTAMP", nil
	}
	return "", nil
}
