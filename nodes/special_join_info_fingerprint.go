// Auto-generated - DO NOT EDIT

package pg_query

import "io"

func (node SpecialJoinInfo) Fingerprint(ctx *FingerprintContext) {
	io.WriteString(ctx.hash, "SpecialJoinInfo")

	for _, subNode := range node.JoinQuals {
		subNode.Fingerprint(ctx)
	}
}
