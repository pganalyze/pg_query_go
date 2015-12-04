// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node FuncCall) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "FUNCCALL")
	io.WriteString(ctx.hash, strconv.FormatBool(node.AggDistinct))

	if node.AggFilter != nil {
		node.AggFilter.Fingerprint(ctx)
	}

	for _, subNode := range node.AggOrder {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.AggStar))
	io.WriteString(ctx.hash, strconv.FormatBool(node.AggWithinGroup))

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.FuncVariadic))

	for _, subNode := range node.Funcname {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Over != nil {
		node.Over.Fingerprint(ctx)
	}
}
