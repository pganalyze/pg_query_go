// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Var) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("Var")
	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.Itoa(int(node.Varattno)))
	ctx.WriteString(strconv.Itoa(int(node.Varcollid)))
	ctx.WriteString(strconv.Itoa(int(node.Varlevelsup)))
	ctx.WriteString(strconv.Itoa(int(node.Varno)))
	ctx.WriteString(strconv.Itoa(int(node.Varnoold)))
	ctx.WriteString(strconv.Itoa(int(node.Varoattno)))
	ctx.WriteString(strconv.Itoa(int(node.Vartype)))
	ctx.WriteString(strconv.Itoa(int(node.Vartypmod)))

	if node.Xpr != nil {
		node.Xpr.Fingerprint(ctx)
	}
}
