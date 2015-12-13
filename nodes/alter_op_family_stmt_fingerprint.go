// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterOpFamilyStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterOpFamilyStmt")

	if node.Amname != nil {
		ctx.WriteString(*node.Amname)
	}

	ctx.WriteString(strconv.FormatBool(node.IsDrop))

	for _, subNode := range node.Items {
		subNode.Fingerprint(ctx, "Items")
	}

	for _, subNode := range node.Opfamilyname {
		subNode.Fingerprint(ctx, "Opfamilyname")
	}
}
