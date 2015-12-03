// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node FuncWithArgs) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FuncWithArgs")

	for _, subNode := range node.Funcargs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}
}
