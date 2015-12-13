// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WindowClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("WindowClause")
	ctx.WriteString(strconv.FormatBool(node.CopiedOrder))

	if node.EndOffset != nil {
		node.EndOffset.Fingerprint(ctx, "EndOffset")
	}

	ctx.WriteString(strconv.Itoa(int(node.FrameOptions)))

	if node.Name != nil {
		ctx.WriteString(*node.Name)
	}

	for _, subNode := range node.OrderClause {
		subNode.Fingerprint(ctx, "OrderClause")
	}

	for _, subNode := range node.PartitionClause {
		subNode.Fingerprint(ctx, "PartitionClause")
	}

	if node.Refname != nil {
		ctx.WriteString(*node.Refname)
	}

	if node.StartOffset != nil {
		node.StartOffset.Fingerprint(ctx, "StartOffset")
	}

	ctx.WriteString(strconv.Itoa(int(node.Winref)))
}
