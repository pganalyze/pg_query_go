// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Join) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("JOIN")

	for _, subNode := range node.Joinqual {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Jointype)))
}
