// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node IndexPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "INDEXPATH")

	for _, subNode := range node.Indexclauses {
		subNode.Fingerprint(ctx)
	}

	if node.Indexinfo != nil {
		node.Indexinfo.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexorderbycols {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexorderbys {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexqualcols {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indexquals {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Indexscandir)))
}
