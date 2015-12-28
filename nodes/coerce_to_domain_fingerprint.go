// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CoerceToDomain) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CoerceToDomain")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if int(node.Coercionformat) != 0 {
		ctx.WriteString("coercionformat")
		ctx.WriteString(strconv.Itoa(int(node.Coercionformat)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Resultcollid != 0 {
		ctx.WriteString("resultcollid")
		ctx.WriteString(strconv.Itoa(int(node.Resultcollid)))
	}

	if node.Resulttype != 0 {
		ctx.WriteString("resulttype")
		ctx.WriteString(strconv.Itoa(int(node.Resulttype)))
	}

	if node.Resulttypmod != 0 {
		ctx.WriteString("resulttypmod")
		ctx.WriteString(strconv.Itoa(int(node.Resulttypmod)))
	}

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
