// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node CreateStmt) Deparse(ctx DeparseContext) (string, error) {
	out := []string{"CREATE"}
	if persistence := node.relPersistence(); persistence != nil {
		out = append(out, *persistence)
	}

	out = append(out, "TABLE")

	if node.IfNotExists {
		out = append(out, "IF NOT EXISTS")
	}

	if str, err := node.Relation.Deparse(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, str)
	}

	elts := make([]string, len(node.TableElts.Items))
	for i, elt := range node.TableElts.Items {
		if str, err := elt.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			elts[i] = str
		}
	}
	out = append(out, fmt.Sprintf("(%s)", strings.Join(elts, ", ")))

	if node.InhRelations.Items != nil && len(node.InhRelations.Items) > 0 {
		out = append(out, "INHERITS")
		relations := make([]string, len(node.InhRelations.Items))
		for i, relation := range node.InhRelations.Items {
			if str, err := relation.Deparse(DeparseContextNone); err != nil {
				return "", err
			} else {
				relations[i] = str
			}
		}
		out = append(out, fmt.Sprintf("(%s)", strings.Join(relations, ", ")))
	}

	return strings.Join(out, " "), nil
}

func (node CreateStmt) relPersistence() *string {
	t, u := "TEMPORARY", "UNLOGGED"
	if string(node.Relation.Relpersistence) == "t" {
		return &t
	} else if string(node.Relation.Relpersistence) == "u" {
		return &u
	}
	return nil
}
