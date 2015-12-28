// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterForeignServerStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterForeignServerStmt")

	if node.HasVersion {
		ctx.WriteString("has_version")
		ctx.WriteString(strconv.FormatBool(node.HasVersion))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Servername != nil {
		ctx.WriteString("servername")
		ctx.WriteString(*node.Servername)
	}

	if node.Version != nil {
		ctx.WriteString("version")
		ctx.WriteString(*node.Version)
	}
}
