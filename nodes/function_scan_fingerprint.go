// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node FunctionScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FunctionScan")

	for _, subNode := range node.Functions {
		subNode.Fingerprint(ctx)
	}
}
