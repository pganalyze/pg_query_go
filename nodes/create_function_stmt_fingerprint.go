// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateFunctionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateFunctionStmt")
	node.Funcname.Fingerprint(ctx, "Funcname")
	node.Options.Fingerprint(ctx, "Options")
	node.Parameters.Fingerprint(ctx, "Parameters")
	ctx.WriteString(strconv.FormatBool(node.Replace))

	if node.ReturnType != nil {
		node.ReturnType.Fingerprint(ctx, "ReturnType")
	}

	node.WithClause.Fingerprint(ctx, "WithClause")
}
