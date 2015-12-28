// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node LockStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("LockStmt")

	if node.Mode != 0 {
		ctx.WriteString("mode")
		ctx.WriteString(strconv.Itoa(int(node.Mode)))
	}

	if node.Nowait {
		ctx.WriteString("nowait")
		ctx.WriteString(strconv.FormatBool(node.Nowait))
	}

	if len(node.Relations.Items) > 0 {
		ctx.WriteString("relations")
		node.Relations.Fingerprint(ctx, "Relations")
	}
}
