// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreatePLangStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreatePLangStmt")

	for _, subNode := range node.Plhandler {
		subNode.Fingerprint(ctx, "Plhandler")
	}

	for _, subNode := range node.Plinline {
		subNode.Fingerprint(ctx, "Plinline")
	}

	if node.Plname != nil {
		ctx.WriteString(*node.Plname)
	}

	ctx.WriteString(strconv.FormatBool(node.Pltrusted))

	for _, subNode := range node.Plvalidator {
		subNode.Fingerprint(ctx, "Plvalidator")
	}

	ctx.WriteString(strconv.FormatBool(node.Replace))
}
