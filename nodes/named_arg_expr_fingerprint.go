// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node NamedArgExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("NamedArgExpr")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx, "Arg")
	}

	ctx.WriteString(strconv.Itoa(int(node.Argnumber)))
	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
