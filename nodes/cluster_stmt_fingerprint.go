// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ClusterStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ClusterStmt")

	if node.Indexname != nil {
		ctx.WriteString("indexname")
		ctx.WriteString(*node.Indexname)
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if node.Verbose {
		ctx.WriteString("verbose")
		ctx.WriteString(strconv.FormatBool(node.Verbose))
	}
}
