// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IntoClause) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("IntoClause")
	if len(node.ColNames.Items) > 0 {
		ctx.WriteString("colNames")
		node.ColNames.Fingerprint(ctx, "ColNames")
	}

	if int(node.OnCommit) != 0 {
		ctx.WriteString("onCommit")
		ctx.WriteString(strconv.Itoa(int(node.OnCommit)))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Rel != nil {
		ctx.WriteString("rel")
		node.Rel.Fingerprint(ctx, "Rel")
	}

	if node.SkipData {
		ctx.WriteString("skipData")
		ctx.WriteString(strconv.FormatBool(node.SkipData))
	}

	if node.TableSpaceName != nil {
		ctx.WriteString("tableSpaceName")
		ctx.WriteString(*node.TableSpaceName)
	}

	if node.ViewQuery != nil {
		ctx.WriteString("viewQuery")
		node.ViewQuery.Fingerprint(ctx, "ViewQuery")
	}
}
