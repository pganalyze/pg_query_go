// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RefreshMatViewStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("RefreshMatViewStmt")
	ctx.WriteString(strconv.FormatBool(node.Concurrent))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.SkipData))
}
