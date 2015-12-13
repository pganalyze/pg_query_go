// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node TableLikeClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("TableLikeClause")
	ctx.WriteString(strconv.Itoa(int(node.Options)))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}
}
