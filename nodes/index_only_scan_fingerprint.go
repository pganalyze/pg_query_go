// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IndexOnlyScan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("INDEXONLYSCAN")

	for _, subNode := range node.Indexorderby {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Indexorderdir)))

	for _, subNode := range node.Indexqual {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indextlist {
		subNode.Fingerprint(ctx)
	}
}
