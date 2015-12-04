// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ParamPathInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "PARAMPATHINFO")

	for _, subNode := range node.PpiClauses {
		subNode.Fingerprint(ctx)
	}
}
