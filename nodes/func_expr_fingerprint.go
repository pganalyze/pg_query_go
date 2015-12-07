// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FuncExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("FUNCEXPR")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Funcformat)))
	ctx.WriteString(strconv.FormatBool(node.Funcretset))
	ctx.WriteString(strconv.FormatBool(node.Funcvariadic))
	// Intentionally ignoring node.Location for fingerprinting
}
