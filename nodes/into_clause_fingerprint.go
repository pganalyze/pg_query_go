// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IntoClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("IntoClause")

	for _, subNode := range node.ColNames {
		subNode.Fingerprint(ctx, "ColNames")
	}

	ctx.WriteString(strconv.Itoa(int(node.OnCommit)))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx, "Options")
	}

	if node.Rel != nil {
		node.Rel.Fingerprint(ctx, "Rel")
	}

	ctx.WriteString(strconv.FormatBool(node.SkipData))

	if node.TableSpaceName != nil {
		ctx.WriteString(*node.TableSpaceName)
	}

	if node.ViewQuery != nil {
		node.ViewQuery.Fingerprint(ctx, "ViewQuery")
	}
}
