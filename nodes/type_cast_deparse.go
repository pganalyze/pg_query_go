// Auto-generated - DO NOT EDIT

package pg_query

import (
	"errors"
	"fmt"
)

func (node TypeCast) Deparse(ctx DeparseContext) (string, error) {
	if node.TypeName == nil {
		return "", errors.New("typename cannot be null in typecast")
	}
	if str, err := node.TypeName.Deparse(DeparseContextNone); err != nil {
		return "", err
	} else {
		if val, err := node.Arg.Deparse(DeparseContextNone); err != nil {
			return "", err
		} else {
			if str == "boolean" {
				if val == "'t'" {
					return "true", nil
				} else {
					return "false", nil
				}
			}

			if typename, err := node.TypeName.Deparse(DeparseContextNone); err != nil {
				return "", err
			} else {
				return fmt.Sprintf("%s::%s", val, typename), nil
			}
		}
	}
}
