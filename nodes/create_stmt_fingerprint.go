// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CreateStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATESTMT")

	for _, subNode := range node.Constraints {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IfNotExists))

	for _, subNode := range node.InhRelations {
		subNode.Fingerprint(ctx)
	}

	if node.OfTypename != nil {
		node.OfTypename.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Oncommit)))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx)
	}

	for _, subNode := range node.TableElts {
		subNode.Fingerprint(ctx)
	}

	if node.Tablespacename != nil {
		io.WriteString(ctx.hash, *node.Tablespacename)
	}
}
