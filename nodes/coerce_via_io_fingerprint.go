// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CoerceViaIO) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CoerceViaIO")

	if node.Arg != nil {
		ctx.WriteString("arg")
		node.Arg.Fingerprint(ctx, "Arg")
	}

	if int(node.Coerceformat) != 0 {
		ctx.WriteString("coerceformat")
		ctx.WriteString(strconv.Itoa(int(node.Coerceformat)))
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

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
