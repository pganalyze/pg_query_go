// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterOwnerStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTEROWNERSTMT")

	if node.Newowner != nil {
		io.WriteString(ctx.hash, *node.Newowner)
	}

	for _, subNode := range node.Objarg {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Object {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.ObjectType)))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}
}
