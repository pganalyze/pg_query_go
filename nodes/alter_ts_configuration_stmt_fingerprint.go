// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node AlterTSConfigurationStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "ALTERTSCONFIGURATIONSTMT")

	for _, subNode := range node.Cfgname {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Dicts {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.MissingOk))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Override))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Replace))

	for _, subNode := range node.Tokentype {
		subNode.Fingerprint(ctx)
	}
}
