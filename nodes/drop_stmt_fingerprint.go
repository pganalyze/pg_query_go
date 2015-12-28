// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node DropStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("DropStmt")
	if len(node.Arguments.Items) > 0 {
		ctx.WriteString("arguments")
		node.Arguments.Fingerprint(ctx, "Arguments")
	}

	if int(node.Behavior) != 0 {
		ctx.WriteString("behavior")
		ctx.WriteString(strconv.Itoa(int(node.Behavior)))
	}

	if node.Concurrent {
		ctx.WriteString("concurrent")
		ctx.WriteString(strconv.FormatBool(node.Concurrent))
	}

	if node.MissingOk {
		ctx.WriteString("missing_ok")
		ctx.WriteString(strconv.FormatBool(node.MissingOk))
	}

	if len(node.Objects.Items) > 0 {
		ctx.WriteString("objects")
		node.Objects.Fingerprint(ctx, "Objects")
	}

	if int(node.RemoveType) != 0 {
		ctx.WriteString("removeType")
		ctx.WriteString(strconv.Itoa(int(node.RemoveType)))
	}
}
