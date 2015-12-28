// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SortBy) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SortBy")
	// Intentionally ignoring node.Location for fingerprinting

	if node.Node != nil {
		ctx.WriteString("node")
		node.Node.Fingerprint(ctx, "Node")
	}

	if int(node.SortbyDir) != 0 {
		ctx.WriteString("sortby_dir")
		ctx.WriteString(strconv.Itoa(int(node.SortbyDir)))
	}

	if int(node.SortbyNulls) != 0 {
		ctx.WriteString("sortby_nulls")
		ctx.WriteString(strconv.Itoa(int(node.SortbyNulls)))
	}

	if len(node.UseOp.Items) > 0 {
		ctx.WriteString("useOp")
		node.UseOp.Fingerprint(ctx, "UseOp")
	}
}
