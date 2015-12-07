// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateOpClassStmt) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("CREATEOPCLASSSTMT")

	if node.Amname != nil {
		ctx.WriteString(*node.Amname)
	}

	if node.Datatype != nil {
		node.Datatype.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.IsDefault))

	for _, subNode := range node.Items {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Opclassname {
		subNode.Fingerprint(ctx)
	}

	for _, subNode := range node.Opfamilyname {
		subNode.Fingerprint(ctx)
	}
}
