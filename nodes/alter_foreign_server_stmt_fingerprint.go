// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterForeignServerStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterForeignServerStmt")
	ctx.WriteString(strconv.FormatBool(node.HasVersion))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}

	if node.Servername != nil {
		ctx.WriteString(*node.Servername)
	}

	if node.Version != nil {
		ctx.WriteString(*node.Version)
	}
}
