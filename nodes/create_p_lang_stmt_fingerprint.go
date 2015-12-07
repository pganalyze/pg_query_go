// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreatePLangStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATEPLANGSTMT")

	for _, subNode := range node.Plhandler {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Plinline {
		subNode.Fingerprint(ctx)
	}

	if node.Plname != nil {
		ctx.WriteString(*node.Plname)
	}

	ctx.WriteString(strconv.FormatBool(node.Pltrusted))

	for _, subNode := range node.Plvalidator {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Replace))
}
