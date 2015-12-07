// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node UniquePath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("UNIQUEPATH")

	for _, subNode := range node.InOperators {
		subNode.Fingerprint(ctx)
	}

	if node.Subpath != nil {
		node.Subpath.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Umethod)))

	for _, subNode := range node.UniqExprs {
		subNode.Fingerprint(ctx)
	}
}
