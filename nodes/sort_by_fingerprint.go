// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SortBy) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SortBy")
	// Intentionally ignoring node.Location for fingerprinting

	if node.Node != nil {
		node.Node.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.SortbyDir)))
	ctx.WriteString(strconv.Itoa(int(node.SortbyNulls)))

	for _, subNode := range node.UseOp {
		subNode.Fingerprint(ctx)
	}
}
