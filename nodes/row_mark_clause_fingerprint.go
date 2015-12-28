// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowMarkClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RowMarkClause")

	if node.NoWait {
		ctx.WriteString("noWait")
		ctx.WriteString(strconv.FormatBool(node.NoWait))
	}

	if node.PushedDown {
		ctx.WriteString("pushedDown")
		ctx.WriteString(strconv.FormatBool(node.PushedDown))
	}

	if node.Rti != 0 {
		ctx.WriteString("rti")
		ctx.WriteString(strconv.Itoa(int(node.Rti)))
	}

	if int(node.Strength) != 0 {
		ctx.WriteString("strength")
		ctx.WriteString(strconv.Itoa(int(node.Strength)))
	}
}
