// Auto-generated - DO NOT EDIT

package pg_query

func (node List) DeparseList(ctx DeparseContext) ([]string, error) {
	out := make([]string, len(node.Items))
	for i, node := range node.Items {
		if str, err := node.Deparse(ctx); err != nil {
			return make([]string, 0), err
		} else {
			out[i] = str
		}
	}
	return out, nil
}

func (node List) Deparse(ctx DeparseContext) (string, error) {
	panic("Not Implemented. Are you sure you didn't mean to use DeparseList()")
}
