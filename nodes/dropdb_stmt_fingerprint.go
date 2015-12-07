// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropdbStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("DROPDBSTMT")

	if node.Dbname != nil {
		ctx.WriteString(*node.Dbname)
	}

	ctx.WriteString(strconv.FormatBool(node.MissingOk))
}
