// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node IndexOnlyScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "IndexOnlyScan")

	for _, subNode := range node.Indexorderby {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexqual {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indextlist {
		subNode.Fingerprint(ctx)
	}
}
