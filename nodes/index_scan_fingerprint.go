// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node IndexScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "IndexScan")

	for _, subNode := range node.Indexorderby {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexorderbyorig {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexqual {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexqualorig {
		subNode.Fingerprint(ctx)
	}
}
