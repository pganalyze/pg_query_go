// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterObjectSchemaStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTEROBJECTSCHEMASTMT")
	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))

	if node.Newschema != nil {
		io.WriteString(ctx.hash, *node.Newschema)
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
