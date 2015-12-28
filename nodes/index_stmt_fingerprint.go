// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IndexStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("IndexStmt")

	if node.AccessMethod != nil {
		ctx.WriteString("accessMethod")
		ctx.WriteString(*node.AccessMethod)
	}

	if node.Concurrent {
		ctx.WriteString("concurrent")
		ctx.WriteString(strconv.FormatBool(node.Concurrent))
	}

	if node.Deferrable {
		ctx.WriteString("deferrable")
		ctx.WriteString(strconv.FormatBool(node.Deferrable))
	}

	if len(node.ExcludeOpNames.Items) > 0 {
		ctx.WriteString("excludeOpNames")
		node.ExcludeOpNames.Fingerprint(ctx, "ExcludeOpNames")
	}

	if node.Idxcomment != nil {
		ctx.WriteString("idxcomment")
		ctx.WriteString(*node.Idxcomment)
	}

	if node.Idxname != nil {
		ctx.WriteString("idxname")
		ctx.WriteString(*node.Idxname)
	}

	if node.IndexOid != 0 {
		ctx.WriteString("indexOid")
		ctx.WriteString(strconv.Itoa(int(node.IndexOid)))
	}

	if len(node.IndexParams.Items) > 0 {
		ctx.WriteString("indexParams")
		node.IndexParams.Fingerprint(ctx, "IndexParams")
	}

	if node.Initdeferred {
		ctx.WriteString("initdeferred")
		ctx.WriteString(strconv.FormatBool(node.Initdeferred))
	}

	if node.Isconstraint {
		ctx.WriteString("isconstraint")
		ctx.WriteString(strconv.FormatBool(node.Isconstraint))
	}

	if node.OldNode != 0 {
		ctx.WriteString("oldNode")
		ctx.WriteString(strconv.Itoa(int(node.OldNode)))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Primary {
		ctx.WriteString("primary")
		ctx.WriteString(strconv.FormatBool(node.Primary))
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if node.TableSpace != nil {
		ctx.WriteString("tableSpace")
		ctx.WriteString(*node.TableSpace)
	}

	if node.Unique {
		ctx.WriteString("unique")
		ctx.WriteString(strconv.FormatBool(node.Unique))
	}

	if node.WhereClause != nil {
		ctx.WriteString("whereClause")
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}
}
