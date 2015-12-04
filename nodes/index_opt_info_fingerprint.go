// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node IndexOptInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "INDEXOPTINFO")
	io.WriteString(ctx.hash, strconv.FormatBool(node.Amcanorderbyop))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Amhasgetbitmap))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Amhasgettuple))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Amoptionalkey))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Amsearcharray))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Amsearchnulls))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Canreturn))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Hypothetical))
	io.WriteString(ctx.hash, strconv.FormatBool(node.Immediate))

	for _, subNode := range node.Indexprs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indextlist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indpred {
		subNode.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.PredOk))

	if node.Rel != nil {
		node.Rel.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Unique))
}
