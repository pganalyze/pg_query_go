// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"fmt"
	"strings"
)

func (node ResTarget) Deparse(ctx DeparseContext) (string, error) {
	switch ctx {
	case DeparseContextNone:
		return *node.Name, nil
	case DeparseContextSelect:
		out := make([]string, 0)
		if str, err := node.Val.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}

		if node.Name != nil && len(*node.Name) > 0 {
			out = append(out, "AS")
			out = append(out, *node.Name)
		}
		return strings.Join(out, " "), nil
	case DeparseContextUpdate:
		out := make([]string, 0)
		if node.Name == nil || len(*node.Name) == 0 {
			return "", errors.New("cannot have blank name for res target in update")
		}
		out = append(out, *node.Name)

		if node.Val == nil {
			return "", errors.New("cannot have null value for res target in update")
		}

		if str, err := node.Val.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			out = append(out, str)
		}

		return strings.Join(out, " = "), nil
	default:
		return "", errors.New(fmt.Sprintf("context type %s is not currently implemented", ctx))
	}
}
