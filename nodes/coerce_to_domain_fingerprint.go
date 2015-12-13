// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CoerceToDomain) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CoerceToDomain")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Coercionformat)))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Resultcollid)))
	ctx.WriteString(strconv.Itoa(int(node.Resulttype)))
	ctx.WriteString(strconv.Itoa(int(node.Resulttypmod)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}
