// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SortBy) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SortBy")
	// Intentionally ignoring node.Location for fingerprinting

	if node.Node != nil {
		node.Node.Fingerprint(ctx, "Node")
	}

	ctx.WriteString(strconv.Itoa(int(node.SortbyDir)))
	ctx.WriteString(strconv.Itoa(int(node.SortbyNulls)))
	node.UseOp.Fingerprint(ctx, "UseOp")
}
