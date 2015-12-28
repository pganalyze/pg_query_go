// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateStmt")
	if len(node.Constraints.Items) > 0 {
		ctx.WriteString("constraints")
		node.Constraints.Fingerprint(ctx, "Constraints")
	}

	if node.IfNotExists {
		ctx.WriteString("if_not_exists")
		ctx.WriteString(strconv.FormatBool(node.IfNotExists))
	}

	if len(node.InhRelations.Items) > 0 {
		ctx.WriteString("inhRelations")
		node.InhRelations.Fingerprint(ctx, "InhRelations")
	}

	if node.OfTypename != nil {
		ctx.WriteString("ofTypename")
		node.OfTypename.Fingerprint(ctx, "OfTypename")
	}

	if int(node.Oncommit) != 0 {
		ctx.WriteString("oncommit")
		ctx.WriteString(strconv.Itoa(int(node.Oncommit)))
	}

	if len(node.Options.Items) > 0 {
		ctx.WriteString("options")
		node.Options.Fingerprint(ctx, "Options")
	}

	if node.Relation != nil {
		ctx.WriteString("relation")
		node.Relation.Fingerprint(ctx, "Relation")
	}

	if len(node.TableElts.Items) > 0 {
		ctx.WriteString("tableElts")
		node.TableElts.Fingerprint(ctx, "TableElts")
	}

	if node.Tablespacename != nil {
		ctx.WriteString("tablespacename")
		ctx.WriteString(*node.Tablespacename)
	}
}
