// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateExtensionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateExtensionStmt")

	if node.Extname != nil {
		ctx.WriteString("extname")
		ctx.WriteString(*node.Extname)
	}

	if node.IfNotExists {
		ctx.WriteString("if_not_exists")
		ctx.WriteString(strconv.FormatBool(node.IfNotExists))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}
}
