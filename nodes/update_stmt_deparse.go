// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"strings"
)

func (node UpdateStmt) Deparse(ctx DeparseContext) (string, error) {
	out := make([]string, 0)

	if node.WithClause != nil {
		if str, err := node.WithClause.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	out = append(out, "UPDATE")

	if node.Relation == nil {
		return "", errors.New("relation of update statement cannot be null")
	}

	if str, err := node.Relation.Deparse(DeparseContextNone); err != nil {
		return "", err
	} else {
		out = append(out, str)
	}

	if node.TargetList.Items == nil || len(node.TargetList.Items) == 0 {
		return "", errors.New("update statement cannot have no sets")
	}

	if node.TargetList.Items != nil && len(node.TargetList.Items) > 0 {
		out = append(out, "SET")
		for _, target := range node.TargetList.Items {
			if str, err := target.Deparse(DeparseContextUpdate); err != nil {
				return "", err
			} else {
				out = append(out, str)
			}
		}
	}

	if node.WhereClause != nil {
		out = append(out, "WHERE")
		if str, err := node.WhereClause.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	if node.ReturningList.Items != nil && len(node.ReturningList.Items) > 0 {
		out = append(out, "RETURNING")
		if returning, err := node.ReturningList.DeparseList(DeparseContextSelect); err != nil {
			return "", err
		} else {
			out = append(out, strings.Join(returning, ", "))
		}
	}
	return strings.Join(out, " "), nil
}
