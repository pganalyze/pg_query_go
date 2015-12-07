// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CoerceToDomain) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("COERCETODOMAIN")

	if node.Arg != nil {
		node.Arg.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Coercionformat)))
	// Intentionally ignoring node.Location for fingerprinting
}
