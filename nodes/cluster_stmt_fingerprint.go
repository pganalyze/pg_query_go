// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ClusterStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ClusterStmt")

	if node.Indexname != nil {
		ctx.WriteString(*node.Indexname)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	ctx.WriteString(strconv.FormatBool(node.Verbose))
}
