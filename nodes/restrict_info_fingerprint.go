// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node RestrictInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "RestrictInfo")
	if node.Clause != nil {
		node.Clause.Fingerprint(ctx)
	}

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

	if node.ParentEc != nil {
		node.ParentEc.Fingerprint(ctx)
	}

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
