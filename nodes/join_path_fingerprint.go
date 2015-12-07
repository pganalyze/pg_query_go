// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node JoinPath) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("JOINPATH")

	if node.Innerjoinpath != nil {
		node.Innerjoinpath.Fingerprint(ctx)
	}

	for _, subNode := range node.Joinrestrictinfo {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Jointype)))

	if node.Outerjoinpath != nil {
		node.Outerjoinpath.Fingerprint(ctx)
	}
}
