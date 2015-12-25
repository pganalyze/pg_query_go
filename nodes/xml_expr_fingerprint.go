// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node XmlExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("XmlExpr")
	node.ArgNames.Fingerprint(ctx, "ArgNames")
	node.Args.Fingerprint(ctx, "Args")
	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	node.NamedArgs.Fingerprint(ctx, "NamedArgs")
	ctx.WriteString(strconv.Itoa(int(node.Op)))
	ctx.WriteString(strconv.Itoa(int(node.Type)))
	ctx.WriteString(strconv.Itoa(int(node.Typmod)))
	ctx.WriteString(strconv.Itoa(int(node.Xmloption)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
