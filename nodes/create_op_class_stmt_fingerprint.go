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

	for _, subNode := range node.Items {
		subNode.Fingerprint(ctx, "Items")
	}

	for _, subNode := range node.Opclassname {
		subNode.Fingerprint(ctx, "Opclassname")
	}

	for _, subNode := range node.Opfamilyname {
		subNode.Fingerprint(ctx, "Opfamilyname")
	}
}
