// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node SecLabelStmt) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SECLABELSTMT")

	if node.Label != nil {
		io.WriteString(ctx.hash, *node.Label)
	}

	for _, subNode := range node.Objargs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Objname {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Objtype)))

	if node.Provider != nil {
		io.WriteString(ctx.hash, *node.Provider)
	}
}
