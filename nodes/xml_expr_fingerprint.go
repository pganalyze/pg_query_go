// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node XmlExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("XmlExpr")
	if len(node.ArgNames.Items) > 0 {
		ctx.WriteString("arg_names")
		node.ArgNames.Fingerprint(ctx, "ArgNames")
	}

	if len(node.Args.Items) > 0 {
		ctx.WriteString("args")
		node.Args.Fingerprint(ctx, "Args")
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}

	if len(node.NamedArgs.Items) > 0 {
		ctx.WriteString("named_args")
		node.NamedArgs.Fingerprint(ctx, "NamedArgs")
	}

	if int(node.Op) != 0 {
		ctx.WriteString("op")
		ctx.WriteString(strconv.Itoa(int(node.Op)))
	}

	if node.Type != 0 {
		ctx.WriteString("type")
		ctx.WriteString(strconv.Itoa(int(node.Type)))
	}

	if node.Typmod != 0 {
		ctx.WriteString("typmod")
		ctx.WriteString(strconv.Itoa(int(node.Typmod)))
	}

	if int(node.Xmloption) != 0 {
		ctx.WriteString("xmloption")
		ctx.WriteString(strconv.Itoa(int(node.Xmloption)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
