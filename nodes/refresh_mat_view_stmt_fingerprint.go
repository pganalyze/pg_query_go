// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RefreshMatViewStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RefreshMatViewStmt")

	if node.Concurrent {
		ctx.WriteString("concurrent")
		ctx.WriteString(strconv.FormatBool(node.Concurrent))
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if node.SkipData {
		ctx.WriteString("skipData")
		ctx.WriteString(strconv.FormatBool(node.SkipData))
	}
}
