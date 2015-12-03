// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node ModifyTable) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ModifyTable")

	for _, subNode := range node.FdwPrivLists {
		subNode.Fingerprint(ctx)
	}

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
