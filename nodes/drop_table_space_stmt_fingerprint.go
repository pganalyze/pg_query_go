// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropTableSpaceStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DropTableSpaceStmt")
	ctx.WriteString(strconv.FormatBool(node.MissingOk))

	if node.Tablespacename != nil {
		ctx.WriteString(*node.Tablespacename)
	}
}
