// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IntoClause) Fingerprint(ctx FingerprintContext) {
	ctx.WriteString("INTOCLAUSE")

	for _, subNode := range node.ColNames {
		subNode.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.Itoa(int(node.OnCommit)))

	for _, subNode := range node.Options {
		subNode.Fingerprint(ctx)
	}

	if node.Rel != nil {
		node.Rel.Fingerprint(ctx)
	}

	ctx.WriteString(strconv.FormatBool(node.SkipData))

	if node.TableSpaceName != nil {
		ctx.WriteString(*node.TableSpaceName)
	}

	if node.ViewQuery != nil {
		node.ViewQuery.Fingerprint(ctx)
	}
}
