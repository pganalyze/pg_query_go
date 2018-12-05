// Auto-generated - DO NOT EDIT

package pg_query

func (node A_Const) Deparse(ctx DeparseContext) (string, error) {
	return node.Val.Deparse(DeparseContextAConst)
}
