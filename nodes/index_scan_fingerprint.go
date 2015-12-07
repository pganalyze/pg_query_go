// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IndexScan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("INDEXSCAN")

	for _, subNode := range node.Indexorderby {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexorderbyorig {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Indexorderdir)))

	for _, subNode := range node.Indexqual {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexqualorig {
		subNode.Fingerprint(ctx)
	}
}
