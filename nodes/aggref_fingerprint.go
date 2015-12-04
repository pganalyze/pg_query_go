// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node Aggref) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "AGGREF")

	for _, subNode := range node.Aggdirectargs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Aggdistinct {
		subNode.Fingerprint(ctx)
	}

	if node.Aggfilter != nil {
		node.Aggfilter.Fingerprint(ctx)
	}

	for _, subNode := range node.Aggorder {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Aggstar))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Aggvariadic))

	for _, subNode := range node.Args {
		subNode.Fingerprint(ctx)
	}

	// Intentionally ignoring node.Location for fingerprinting
}
