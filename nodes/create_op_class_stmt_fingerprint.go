// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateOpClassStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateOpClassStmt")

	if node.Amname != nil {
		ctx.WriteString(*node.Amname)
	}

	if node.Datatype != nil {
		node.Datatype.Fingerprint(ctx, "Datatype")
	}

	ctx.WriteString(strconv.FormatBool(node.IsDefault))
	node.Items.Fingerprint(ctx, "Items")
	node.Opclassname.Fingerprint(ctx, "Opclassname")
	node.Opfamilyname.Fingerprint(ctx, "Opfamilyname")
}
