// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node ForeignScan) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("FOREIGNSCAN")

	for _, subNode := range node.FdwExprs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.FdwPrivate {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.FsSystemCol))
}
