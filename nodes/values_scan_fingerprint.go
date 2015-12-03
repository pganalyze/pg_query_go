// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ValuesScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ValuesScan")

	for _, subNode := range node.ValuesLists {
		subNode.Fingerprint(ctx)
	}
}
