// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node IndexOnlyScan) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "INDEXONLYSCAN")

	for _, subNode := range node.Indexorderby {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Indexorderdir)))

	for _, subNode := range node.Indexqual {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indextlist {
		subNode.Fingerprint(ctx)
	}
}
