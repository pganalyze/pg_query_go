// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ModifyTable) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("MODIFYTABLE")
	ctx.WriteString(strconv.FormatBool(node.CanSetTag))

	for _, subNode := range node.FdwPrivLists {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Operation)))

	for _, subNode := range node.Plans {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ResultRelations {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.ReturningLists {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.RowMarks {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.WithCheckOptionLists {
		subNode.Fingerprint(ctx)
	}
}
