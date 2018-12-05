// Auto-generated - DO NOT EDIT

package pg_query

import (
	"fmt"
	"strings"
)

func (node SelectStmt) Deparse(ctx DeparseContext) (string, error) {
	out := make([]string, 0)
	if node.Op == SETOP_UNION {
		if str, err := node.Larg.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}

		out = append(out, "UNION")
		if node.All {
			out = append(out, "ALL")
		}

		if str, err := node.Rarg.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}

		return strings.Join(out, " "), nil
	}

	if node.WithClause != nil {
		if str, err := node.WithClause.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	// Get select *distinct* *fields*
	if node.TargetList.Items != nil && len(node.TargetList.Items) > 0 {
		out = append(out, "SELECT")
		if node.DistinctClause.Items != nil && len(node.DistinctClause.Items) > 0 {
			out = append(out, "DISTINCT")
		}
		if fields, err := node.TargetList.DeparseList(DeparseContextSelect); err != nil {
			return "", err
		} else {
			out = append(out, strings.Join(fields, ", "))
		}
	}

	if node.FromClause.Items != nil && len(node.FromClause.Items) > 0 {
		out = append(out, "FROM")
		if froms, err := node.FromClause.DeparseList(DeparseContextSelect); err != nil {
			return "", err
		} else {
			out = append(out, strings.Join(froms, ", "))
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

	if node.ValuesLists != nil && len(node.ValuesLists) > 0 {
		out = append(out, "VALUES")
		for _, valueList := range node.ValuesLists {
			values := make([]string, len(valueList))
			for i, value := range valueList {
				if str, err := value.Deparse(DeparseContextNone); err != nil {
					return "", err
				} else {
					values[i] = str
				}
			}
			out = append(out, fmt.Sprintf("(%s)", strings.Join(values, ", ")))
		}
	}

	if node.GroupClause.Items != nil && len(node.GroupClause.Items) > 0 {
		out = append(out, "GROUP BY")
		if groups, err := node.GroupClause.DeparseList(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, strings.Join(groups, ", "))
		}
	}

	if node.HavingClause != nil {
		if str, err := node.HavingClause.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	// Sort clause

	if node.LimitCount != nil {
		out = append(out, "LIMIT")
		if str, err := node.LimitCount.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	if node.LimitOffset != nil {
		out = append(out, "OFFSET")
		if str, err := node.LimitOffset.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}
	}

	if node.LockingClause.Items != nil && len(node.LockingClause.Items) > 0 {
		for _, lock := range node.LockingClause.Items {
			if str, err := lock.Deparse(DeparseContextNone); err != nil {
				return "", err
			} else {
				out = append(out, str)
			}
		}
	}

	return strings.Join(out, " "), nil
}
