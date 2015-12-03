// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node AppendRelInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AppendRelInfo")

	for _, subNode := range node.TranslatedVars {
		subNode.Fingerprint(ctx)
	}
}
