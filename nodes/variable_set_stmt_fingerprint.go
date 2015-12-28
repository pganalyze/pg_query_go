// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node VariableSetStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("VariableSetStmt")
	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	if node.IsLocal {
		ctx.WriteString("is_local")
		ctx.WriteString(strconv.FormatBool(node.IsLocal))
	}

	if int(node.Kind) != 0 {
		ctx.WriteString("kind")
		ctx.WriteString(strconv.Itoa(int(node.Kind)))
	}

	if node.Name != nil {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}
}
