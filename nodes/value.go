package pg_query

import (
  "encoding/json"
  "fmt"
)

/*----------------------
 *		Value node
 *
 * The same Value struct is used for five node types: T_Integer,
 * T_Float, T_String, T_BitString, T_Null.
 *
 * Integral values are actually represented by a machine integer,
 * but both floats and strings are represented as strings.
 * Using T_Float as the node type simply indicates that
 * the contents of the string look like a valid numeric literal.
 *
 * (Before Postgres 7.0, we used a double to represent T_Float,
 * but that creates loss-of-precision problems when the value is
 * ultimately destined to be converted to NUMERIC.  Since Value nodes
 * are only used in the parsing process, not for runtime data, it's
 * better to use the more general representation.)
 *
 * Note that an integer-looking string will get lexed as T_Float if
 * the value is too large to fit in a 'long'.
 *
 * Nulls, of course, don't need the value part at all.
 *----------------------
 */

type Value struct {
  Type NodeTag			/* tag appropriately (eg. T_String) */
  Ival int `json:"ival"`		/* machine integer */
  Str string `json:"str"`		/* string */
}

func (input Value) MarshalJSON() ([]byte, error) {
  switch input.Type {
  case T_Integer:
    return json.Marshal(input.Ival)
  case T_Float:
    return json.Marshal(input.Str)
  case T_String:
    return json.Marshal(input.Str)
  case T_BitString:
    return json.Marshal(input.Str)
  case T_Null:
    return json.Marshal(nil)
  }

  return json.Marshal(nil)
}

func (value *Value) UnmarshalJSON(input []byte) (err error) {
  var i interface{}
  err = json.Unmarshal(input, &i)
  if err != nil {
    return
  }

  switch v := i.(type) {
  case float64:
    value.Type = T_Integer // FIXME: Support floats (check if input has a dot in it?)
    value.Ival = int(v)
  case string:
    value.Type = T_String
    value.Str = v
  case nil:
    value.Type = T_Null
  default:
    err = fmt.Errorf("Unsupported value %V", v)
  }

  return
}

func (value Value) Deparse() string {
  panic("Not Implemented")
}
