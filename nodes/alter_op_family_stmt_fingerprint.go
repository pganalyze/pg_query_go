// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterOpFamilyStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterOpFamilyStmt")

	if node.Amname != nil {
		ctx.WriteString("amname")
		ctx.WriteString(*node.Amname)
	}

	if node.IsDrop {
		ctx.WriteString("isDrop")
		ctx.WriteString(strconv.FormatBool(node.IsDrop))
	}

	if len(node.Items.Items) > 0 {
		ctx.WriteString("items")
		node.Items.Fingerprint(ctx, "Items")
	}

	if len(node.Opfamilyname.Items) > 0 {
		ctx.WriteString("opfamilyname")
		node.Opfamilyname.Fingerprint(ctx, "Opfamilyname")
	}
}
