// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CopyStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CopyStmt")
	node.Attlist.Fingerprint(ctx, "Attlist")

	if node.Filename != nil {
		ctx.WriteString(*node.Filename)
	}

	ctx.WriteString(strconv.FormatBool(node.IsFrom))
	ctx.WriteString(strconv.FormatBool(node.IsProgram))
	node.Options.Fingerprint(ctx, "Options")

	if node.Query != nil {
		node.Query.Fingerprint(ctx, "Query")
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}
}
