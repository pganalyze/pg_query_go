// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node OpExpr) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("OPEXPR")

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	ctx.WriteString(strconv.FormatBool(node.Opretset))
}
