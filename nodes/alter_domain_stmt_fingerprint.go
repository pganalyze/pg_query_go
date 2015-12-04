// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterDomainStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERDOMAINSTMT")
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Behavior)))

	if node.Def != nil {
		node.Def.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}

	for _, subNode := range node.TypeName {
		subNode.Fingerprint(ctx)
	}
}
