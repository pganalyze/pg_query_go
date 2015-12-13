// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FetchStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FetchStmt")
	ctx.WriteString(strconv.Itoa(int(node.Direction)))
	ctx.WriteString(strconv.Itoa(int(node.HowMany)))
	ctx.WriteString(strconv.FormatBool(node.Ismove))

	if node.Portalname != nil {
		ctx.WriteString(*node.Portalname)
	}
}
