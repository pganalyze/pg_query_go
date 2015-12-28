// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CopyStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CopyStmt")
	if len(node.Attlist.Items) > 0 {
		ctx.WriteString("attlist")
		node.Attlist.Fingerprint(ctx, "Attlist")
	}

	if node.Filename != nil {
		ctx.WriteString("filename")
		ctx.WriteString(*node.Filename)
	}

	if node.IsFrom {
		ctx.WriteString("is_from")
		ctx.WriteString(strconv.FormatBool(node.IsFrom))
	}

	if node.IsProgram {
		ctx.WriteString("is_program")
		ctx.WriteString(strconv.FormatBool(node.IsProgram))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Query != nil {
		ctx.WriteString("query")
		node.Query.Fingerprint(ctx, "Query")
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}
}
