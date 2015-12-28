// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateOpClassStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateOpClassStmt")

	if node.Amname != nil {
		ctx.WriteString("amname")
		ctx.WriteString(*node.Amname)
	}

	if node.Datatype != nil {
		ctx.WriteString("datatype")
		node.Datatype.Fingerprint(ctx, "Datatype")
	}

	if node.IsDefault {
		ctx.WriteString("isDefault")
		ctx.WriteString(strconv.FormatBool(node.IsDefault))
	}

	if len(node.Items.Items) > 0 {
		ctx.WriteString("items")
		node.Items.Fingerprint(ctx, "Items")
	}

	if len(node.Opclassname.Items) > 0 {
		ctx.WriteString("opclassname")
		node.Opclassname.Fingerprint(ctx, "Opclassname")
	}

	if len(node.Opfamilyname.Items) > 0 {
		ctx.WriteString("opfamilyname")
		node.Opfamilyname.Fingerprint(ctx, "Opfamilyname")
	}
}
