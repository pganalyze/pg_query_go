// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node XmlExpr) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "XMLEXPR")

	for _, subNode := range node.ArgNames {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Name != nil {
		io.WriteString(ctx.hash, *node.Name)
	}

	for _, subNode := range node.NamedArgs {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.Itoa(int(node.Op)))
	io.WriteString(ctx.hash, strconv.Itoa(int(node.Xmloption)))
}
