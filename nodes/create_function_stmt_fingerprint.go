// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateFunctionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateFunctionStmt")
	if len(node.Funcname.Items) > 0 {
		ctx.WriteString("funcname")
		node.Funcname.Fingerprint(ctx, "Funcname")
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if len(node.Parameters.Items) > 0 {
		ctx.WriteString("parameters")
		node.Parameters.Fingerprint(ctx, "Parameters")
	}

	if node.Replace {
		ctx.WriteString("replace")
		ctx.WriteString(strconv.FormatBool(node.Replace))
	}

	if node.ReturnType != nil {
		ctx.WriteString("returnType")
		node.ReturnType.Fingerprint(ctx, "ReturnType")
	}

	if len(node.WithClause.Items) > 0 {
		ctx.WriteString("withClause")
		node.WithClause.Fingerprint(ctx, "WithClause")
	}
}
