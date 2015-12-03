// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node WindowFunc) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "WindowFunc")
	if node.Aggfilter != nil {
		node.Aggfilter.Fingerprint(ctx)
	}

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}
}
