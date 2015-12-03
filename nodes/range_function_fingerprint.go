// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RangeFunction) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RangeFunction")
	if node.Alias != nil {
		node.Alias.Fingerprint(ctx)
	}

	for _, subNode := range node.Coldeflist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Functions {
		subNode.Fingerprint(ctx)
	}
}
