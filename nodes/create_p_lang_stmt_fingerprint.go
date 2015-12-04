// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node CreatePLangStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "CREATEPLANGSTMT")

	for _, subNode := range node.Plhandler {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Plinline {
		subNode.Fingerprint(ctx)
	}

	if node.Plname != nil {
		io.WriteString(ctx.hash, *node.Plname)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Pltrusted))

	for _, subNode := range node.Plvalidator {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Replace))
}
