// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node GrantStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "GRANTSTMT")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Behavior)))
	io.WriteString(ctx.hash, strconv.FormatBool(node.GrantOption))

	for _, subNode := range node.Grantees {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IsGrant))

	for _, subNode := range node.Objects {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Objtype)))

	for _, subNode := range node.Privileges {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Targtype)))
}
