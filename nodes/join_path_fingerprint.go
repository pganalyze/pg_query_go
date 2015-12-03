// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node JoinPath) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "JoinPath")
	if node.Innerjoinpath != nil {
		node.Innerjoinpath.Fingerprint(ctx)
	}

	for _, subNode := range node.Joinrestrictinfo {
		subNode.Fingerprint(ctx)
	}

	if node.Outerjoinpath != nil {
		node.Outerjoinpath.Fingerprint(ctx)
	}
}
