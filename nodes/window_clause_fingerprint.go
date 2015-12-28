// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node WindowClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("WindowClause")

	if node.CopiedOrder {
		ctx.WriteString("copiedOrder")
		ctx.WriteString(strconv.FormatBool(node.CopiedOrder))
	}

	if node.EndOffset != nil {
		ctx.WriteString("endOffset")
		node.EndOffset.Fingerprint(ctx, "EndOffset")
	}

	if node.FrameOptions != 0 {
		ctx.WriteString("frameOptions")
		ctx.WriteString(strconv.Itoa(int(node.FrameOptions)))
	}

	if node.Name != nil {
		ctx.WriteString("name")
		ctx.WriteString(*node.Name)
	}

	if len(node.OrderClause.Items) > 0 {
		ctx.WriteString("orderClause")
		node.OrderClause.Fingerprint(ctx, "OrderClause")
	}

	if len(node.PartitionClause.Items) > 0 {
		ctx.WriteString("partitionClause")
		node.PartitionClause.Fingerprint(ctx, "PartitionClause")
	}

	if node.Refname != nil {
		ctx.WriteString("refname")
		ctx.WriteString(*node.Refname)
	}

	if node.StartOffset != nil {
		ctx.WriteString("startOffset")
		node.StartOffset.Fingerprint(ctx, "StartOffset")
	}

	if node.Winref != 0 {
		ctx.WriteString("winref")
		ctx.WriteString(strconv.Itoa(int(node.Winref)))
	}
}
