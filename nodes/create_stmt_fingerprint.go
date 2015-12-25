// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateStmt")
	node.Constraints.Fingerprint(ctx, "Constraints")
	ctx.WriteString(strconv.FormatBool(node.IfNotExists))
	node.InhRelations.Fingerprint(ctx, "InhRelations")

	if node.OfTypename != nil {
		node.OfTypename.Fingerprint(ctx, "OfTypename")
	}

	ctx.WriteString(strconv.Itoa(int(node.Oncommit)))
	node.Options.Fingerprint(ctx, "Options")

	if node.Relation != nil {
		node.Relation.Fingerprint(ctx, "Relation")
	}

	node.TableElts.Fingerprint(ctx, "TableElts")

	if node.Tablespacename != nil {
		ctx.WriteString(*node.Tablespacename)
	}
}
