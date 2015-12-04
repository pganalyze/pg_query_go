// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node ModifyTable) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "MODIFYTABLE")
	io.WriteString(ctx.hash, strconv.FormatBool(node.CanSetTag))

	for _, subNode := range node.FdwPrivLists {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Operation)))

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
