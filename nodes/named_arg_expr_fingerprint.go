// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node NamedArgExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("NamedArgExpr")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if node.Argnumber != 0 {
		ctx.WriteString("argnumber")
		ctx.WriteString(strconv.Itoa(int(node.Argnumber)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
