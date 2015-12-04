// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RenameStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RENAMESTMT")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Behavior)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))

	if node.Newname != nil {
		io.WriteString(ctx.hash, *node.Newname)
	}

	for _, subNode := range node.Objarg {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Object {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.RelationType)))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.RenameType)))

	if node.Subname != nil {
		io.WriteString(ctx.hash, *node.Subname)
	}
}
