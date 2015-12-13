// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CoerceViaIO) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CoerceViaIO")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Coerceformat)))
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Resultcollid)))
	ctx.WriteString(strconv.Itoa(int(node.Resulttype)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}
