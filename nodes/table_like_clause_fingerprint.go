// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TableLikeClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TableLikeClause")

	if node.Options != 0 {
		ctx.WriteString("options")
		ctx.WriteString(strconv.Itoa(int(node.Options)))
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}
}
