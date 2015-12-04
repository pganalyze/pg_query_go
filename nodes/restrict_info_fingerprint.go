// Auto-generated - DO NOT EDIT

package pg_query

import (
	"io"
	"strconv"
)

func (node RestrictInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RESTRICTINFO")
	io.WriteString(ctx.hash, strconv.FormatBool(node.CanJoin))

	if node.Clause != nil {
		node.Clause.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.IsPushedDown))

	if node.LeftEc != nil {
		node.LeftEc.Fingerprint(ctx)
	}

	if node.LeftEm != nil {
		node.LeftEm.Fingerprint(ctx)
	}

	for _, subNode := range node.Mergeopfamilies {
		subNode.Fingerprint(ctx)
	}

	if node.Orclause != nil {
		node.Orclause.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.OuterIsLeft))
	io.WriteString(ctx.hash, strconv.FormatBool(node.OuterjoinDelayed))

	if node.ParentEc != nil {
		node.ParentEc.Fingerprint(ctx)
	}

	io.WriteString(ctx.hash, strconv.FormatBool(node.Pseudoconstant))

	if node.RightEc != nil {
		node.RightEc.Fingerprint(ctx)
	}

	if node.RightEm != nil {
		node.RightEm.Fingerprint(ctx)
	}

	for _, subNode := range node.ScanselCache {
		subNode.Fingerprint(ctx)
	}
}
