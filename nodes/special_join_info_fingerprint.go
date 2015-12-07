// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SpecialJoinInfo) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("SPECIALJOININFO")
	ctx.WriteString(strconv.FormatBool(node.DelayUpperJoins))

	for _, subNode := range node.JoinQuals {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.Jointype)))
	ctx.WriteString(strconv.FormatBool(node.LhsStrict))
}
