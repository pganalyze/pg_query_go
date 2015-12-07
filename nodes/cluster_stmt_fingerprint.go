// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ClusterStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CLUSTERSTMT")

	if node.Indexname != nil {
		ctx.WriteString(*node.Indexname)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Verbose))
}
