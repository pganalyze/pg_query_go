// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterEnumStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterEnumStmt")

	if node.NewVal != nil {
		ctx.WriteString(*node.NewVal)
	}

	ctx.WriteString(strconv.FormatBool(node.NewValIsAfter))

	if node.NewValNeighbor != nil {
		ctx.WriteString(*node.NewValNeighbor)
	}

	ctx.WriteString(strconv.FormatBool(node.SkipIfExists))

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx, "TypeName")
	}
}
