// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node IndexStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("IndexStmt")

	if node.AccessMethod != nil {
		ctx.WriteString(*node.AccessMethod)
	}

	ctx.WriteString(strconv.FormatBool(node.Concurrent))
	ctx.WriteString(strconv.FormatBool(node.Deferrable))
	node.ExcludeOpNames.Fingerprint(ctx, "ExcludeOpNames")

	if node.Idxcomment != nil {
		ctx.WriteString(*node.Idxcomment)
	}

	if node.Idxname != nil {
		ctx.WriteString(*node.Idxname)
	}

	ctx.WriteString(strconv.Itoa(int(node.IndexOid)))
	node.IndexParams.Fingerprint(ctx, "IndexParams")
	ctx.WriteString(strconv.FormatBool(node.Initdeferred))
	ctx.WriteString(strconv.FormatBool(node.Isconstraint))
	ctx.WriteString(strconv.Itoa(int(node.OldNode)))
	node.Options.Fingerprint(ctx, "Options")
	ctx.WriteString(strconv.FormatBool(node.Primary))

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if node.TableSpace != nil {
		ctx.WriteString(*node.TableSpace)
	}

	ctx.WriteString(strconv.FormatBool(node.Unique))

	if node.WhereClause != nil {
		node.WhereClause.Fingerprint(ctx, "WhereClause")
	}
}
