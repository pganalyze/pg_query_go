// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IndexOptInfo) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("INDEXOPTINFO")
	ctx.WriteString(strconv.FormatBool(node.Amcanorderbyop))
	ctx.WriteString(strconv.FormatBool(node.Amhasgetbitmap))
	ctx.WriteString(strconv.FormatBool(node.Amhasgettuple))
	ctx.WriteString(strconv.FormatBool(node.Amoptionalkey))
	ctx.WriteString(strconv.FormatBool(node.Amsearcharray))
	ctx.WriteString(strconv.FormatBool(node.Amsearchnulls))
	ctx.WriteString(strconv.FormatBool(node.Canreturn))
	ctx.WriteString(strconv.FormatBool(node.Hypothetical))
	ctx.WriteString(strconv.FormatBool(node.Immediate))

	for _, subNode := range node.Indexprs {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indextlist {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Indpred {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.PredOk))

	if node.Rel != nil {
		node.Rel.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.Unique))
}
