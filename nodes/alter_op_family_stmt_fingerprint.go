// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterOpFamilyStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("ALTEROPFAMILYSTMT")

	if node.Amname != nil {
		ctx.WriteString(*node.Amname)
	}

	ctx.WriteString(strconv.FormatBool(node.IsDrop))

	for _, subNode := range node.Items {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Opfamilyname {
		subNode.Fingerprint(ctx)
	}
}
