package pg_query

import (
	"fmt"
	"log"
	reflect "reflect"
	"strings"
)

var debugLog bool

type TagType int

const (
	TagDDL TagType = iota
	TagDML
	TagSelect
	TagCall
	TagFuncArg
)

type TaggedNode struct {
	Index int
	Item  isNode_Node
	Name  string
	T     TagType
}

type TaggedParseResult struct {
	pr            *ParseResult
	tables        []TaggedNode
	functions     []TaggedNode
	functionsArgs []TaggedNode
	Aliases       map[string]string
	CteNames      map[string]bool
}

type Column struct {
	Schema *string
	Rel    *string
	Index  int
}

func (c *Column) String() (s string) {
	s += fmt.Sprintf("{%d, ", c.Index)
	if c.Schema == nil {
		s += "nil, "
	} else {
		s += fmt.Sprintf("%s, ", *c.Schema)
	}
	if c.Rel == nil {
		s += "nil}"
	} else {
		s += fmt.Sprintf("%s}", *c.Rel)
	}
	return s
}

// FilterColumns is a function that returns a list of columns that the query
// filters by except the target list, but includes things like JOIN condition
// and WHERE clause and traverses sub-selects.
func (tpr *TaggedParseResult) FilterColumns() (filtersColumns []Column, err error) {
	// Get condition items from the parse tree
	var stms []*Node
	var conditions []*Node
	for _, e := range tpr.pr.Stmts {
		if e == nil {
			continue
		}
		stms = append(stms, e.GetStmt())
	}
	for len(stms) > 0 || len(conditions) > 0 {
		if len(stms) > 0 {
			e := stms[:1][0].GetNode()
			stms = stms[1:]
			if e == nil {
				continue
			}
			switch node := e.(type) {
			//  *Node_List
			case *Node_List:
				// append list to stms
				stms = append(stms, node.List.Items...)
			//	*Node_RawStmt
			case *Node_RawStmt:
				stms = append(stms, node.RawStmt.Stmt)
			//	*Node_SelectStmt
			case *Node_SelectStmt:
				switch node.SelectStmt.Op {
				case SetOperation_SETOP_NONE:
					if len(node.SelectStmt.FromClause) > 0 {
						// From sub selects
						for _, e := range node.SelectStmt.FromClause {
							node := e.GetNode()
							rangeSubselect, ok := node.(*Node_RangeSubselect)
							if ok {
								stms = append(stms, rangeSubselect.RangeSubselect.Subquery)
							}
						}
						// Join on conditions
						conditions = append(conditions, conditionsFromJoinClauses(node.SelectStmt.FromClause)...)

						// Where clause
						if node.SelectStmt.WhereClause != nil {
							conditions = append(conditions, node.SelectStmt.WhereClause)
						}

						// With Clause
						if node.SelectStmt.WithClause != nil {
							for _, e := range node.SelectStmt.WithClause.Ctes {
								n := e.GetNode().(*Node_CommonTableExpr)
								stms = append(stms, n.CommonTableExpr.Ctequery)
							}
						}
					}
				case SetOperation_SETOP_UNION, SetOperation_SETOP_INTERSECT, SetOperation_SETOP_EXCEPT:
					if node.SelectStmt.Larg != nil {
						larg := node.SelectStmt.Larg
						stms = append(stms, &Node{Node: &Node_SelectStmt{SelectStmt: larg}})
					}
					if node.SelectStmt.Rarg != nil {
						rarg := node.SelectStmt.Rarg
						stms = append(stms, &Node{Node: &Node_SelectStmt{SelectStmt: rarg}})
					}
				}
			//	*Node_UpdateStmt
			case *Node_UpdateStmt:
				if node.UpdateStmt.WhereClause != nil {
					conditions = append(conditions, node.UpdateStmt.WhereClause)
				}
			//	*Node_DeleteStmt
			case *Node_DeleteStmt:
				if node.DeleteStmt.WhereClause != nil {
					conditions = append(conditions, node.DeleteStmt.WhereClause)
				}
			//	*Node_IndexStmt
			case *Node_IndexStmt:
				if node.IndexStmt.WhereClause != nil {
					conditions = append(conditions, node.IndexStmt.WhereClause)
				}
			}
		}

		// Process both JOIN and WHERE conditions here
		if len(conditions) > 0 {
			if conditions[0] == nil {
				conditions = conditions[1:]
				continue
			}
			node := conditions[:1][0].GetNode()
			conditions = conditions[1:]
			switch node := node.(type) {
			case *Node_AExpr:
				if node.AExpr.Lexpr != nil {
					lexpr := node.AExpr.Lexpr
					conditions = append(conditions, lexpr)
				}
				if node.AExpr.Rexpr != nil {
					rexpr := node.AExpr.Rexpr
					conditions = append(conditions, rexpr)
				}
			case *Node_BoolExpr:
				conditions = append(conditions, node.BoolExpr.Args...)
			case *Node_CoalesceExpr:
				conditions = append(conditions, node.CoalesceExpr.Args...)
			case *Node_RowExpr:
				conditions = append(conditions, node.RowExpr.Args...)
			case *Node_ColumnRef:
				var strbuf []string
				var indexes []int
				for index, e := range node.ColumnRef.Fields {
					strbuf = append(strbuf, e.GetNode().(*Node_String_).String_.GetSval())
					indexes = append(indexes, index)
				}

				if len(strbuf) == 1 {
					filtersColumns = append(filtersColumns, Column{Schema: nil, Rel: &strbuf[0]})
				} else if len(strbuf) == 2 {
					if a, ok := tpr.Aliases[strbuf[0]]; ok {
						filtersColumns = append(filtersColumns, Column{Schema: &a, Rel: &strbuf[1]})
					} else {
						filtersColumns = append(filtersColumns, Column{Schema: &strbuf[0], Rel: &strbuf[1]})
					}
				} else {
					return nil, fmt.Errorf("ColumnRef doesn't have 1-2 fields, got %d", len(strbuf))
				}
			case *Node_NullTest:
				conditions = append(conditions, node.NullTest.Arg)
			case *Node_BooleanTest:
				conditions = append(conditions, node.BooleanTest.Arg)
			case *Node_FuncCall:
				//FIXME: This should actually be extracted as a funccall and be compared with those indices
				if len(node.FuncCall.Args) > 0 {
					conditions = append(conditions, node.FuncCall.Args...)
				}
			case *Node_SubLink:
				conditions = append(conditions, node.SubLink.Testexpr)
				stms = append(stms, node.SubLink.Subselect)
			}
		}
	}

	// TODO Make pairs uniq
	return filtersColumns, nil
}

func conditionsFromJoinClauses(joinClauses []*Node) (conditions []*Node) {
	for _, e := range joinClauses {
		if e == nil {
			continue
		}
		_, ok := e.GetNode().(*Node_JoinExpr)
		if !ok {
			continue
		}
		var joinexprItems []*Node
		joinexprItems = append(joinexprItems, e)
		for len(joinexprItems) > 0 {
			e := joinexprItems[:1][0].GetNode().(*Node_JoinExpr)
			joinexprItems = joinexprItems[1:]
			if e.JoinExpr.Quals != nil {
				conditions = append(conditions, e.JoinExpr.Quals)
			}
			if e.JoinExpr.Larg.GetJoinExpr() != nil {
				joinexprItems = append(joinexprItems, e.JoinExpr.Larg)
			}
			if e.JoinExpr.Rarg.GetJoinExpr() != nil {
				joinexprItems = append(joinexprItems, e.JoinExpr.Rarg)
			}
		}
	}
	return conditions
}

func statements_and_cte_names_for_with_clause(withClause *WithClause) (stms []*Node, cteNames []string) {
	if withClause != nil {
		for _, e := range withClause.Ctes {
			if e == nil {
				continue
			}
			cte := e.GetCommonTableExpr()
			stms = append(stms, cte.Ctequery)
			cteNames = append(cteNames, cte.Ctename)
		}
	}
	return stms, cteNames
}

// Getters

type TaggedTableFilter func(*TaggedNode) bool

// FilterTables filters the tables list and returns the filtered list
func (tpr *TaggedParseResult) filterNode(p TaggedTableFilter, nodes []TaggedNode) []TaggedNode {
	// It uses the tables list and filters out for similar names
	objNames := map[string]bool{}
	var filteredObjects []TaggedNode
	for _, e := range nodes {
		if !p(&e) || objNames[e.Name] {
			continue
		}
		filteredObjects = append(filteredObjects, e)
		objNames[e.Name] = true
	}
	return filteredObjects
}

// GetTables returns the tables list
func (tpr *TaggedParseResult) GetTables() []TaggedNode {
	return tpr.filterNode(func(t *TaggedNode) bool {
		return true
	}, tpr.tables)
}

func (tpr *TaggedParseResult) GetSelectTables() []TaggedNode {
	return tpr.filterNode(func(t *TaggedNode) bool {
		return t.T == TagSelect
	}, tpr.tables)
}

func (tpr *TaggedParseResult) GetDMLTables() []TaggedNode {
	return tpr.filterNode(func(t *TaggedNode) bool {
		return t.T == TagDML
	}, tpr.tables)
}

func (tpr *TaggedParseResult) GetDDLTables() []TaggedNode {
	return tpr.filterNode(func(t *TaggedNode) bool {
		return t.T == TagDDL
	}, tpr.tables)
}

func (tpr *TaggedParseResult) GetFunctions() []TaggedNode {
	return tpr.filterNode(func(t *TaggedNode) bool {
		return true
	}, tpr.functions)
}

func (tpr *TaggedParseResult) GetCallFunctions() []TaggedNode {
	return tpr.filterNode(func(t *TaggedNode) bool {
		return t.T == TagCall
	}, tpr.functions)
}

func (tpr *TaggedParseResult) GetDDLFunctions() []TaggedNode {
	return tpr.filterNode(func(t *TaggedNode) bool {
		return t.T == TagDDL
	}, tpr.functions)
}

func (tpr *TaggedParseResult) GetArgsFunctions() []TaggedNode {
	return tpr.filterNode(func(t *TaggedNode) bool {
		return t.T == TagFuncArg
	}, tpr.functionsArgs)
}

// NewTaggedParseResult creates a new TaggedParseResult using the given ParseResult
func NewTaggedParseResult(pr *ParseResult) (tpr *TaggedParseResult) {
	tpr = &TaggedParseResult{pr: pr, Aliases: map[string]string{}, CteNames: map[string]bool{}}
	return tpr
}

// LoadObjects Parses the query and finds table and function references
func (tpr *TaggedParseResult) LoadObjects() (err error) {
	var stms []*Node
	var subSelectItems []*Node
	var fromClause []TaggedNode

	// map from RawStmt to *Node
	for _, e := range tpr.pr.Stmts {
		if e == nil {
			continue
		}
		stms = append(stms, e.GetStmt())
	}
	for len(stms) > 0 || len(fromClause) > 0 || len(subSelectItems) > 0 {
		if len(stms) > 0 {
			e := stms[:1][0].GetNode()
			stms = stms[1:]
			switch node := e.(type) {
			case *Node_List:
				stms = append(stms, node.List.Items...)
			// The following statement types do not modify tables and are added to fromClause
			// and subsequently to tables variable
			case *Node_SelectStmt:
				subSelectItems = append(subSelectItems, node.SelectStmt.TargetList...)
				if node.SelectStmt.WhereClause != nil {
					subSelectItems = append(subSelectItems, node.SelectStmt.WhereClause)
				}
				var buf []*Node
				for _, e := range node.SelectStmt.SortClause {
					buf = append(buf, e.GetSortBy().GetNode())
				}
				subSelectItems = append(subSelectItems, buf...)
				subSelectItems = append(subSelectItems, node.SelectStmt.GroupClause...)
				if node.SelectStmt.HavingClause != nil {
					subSelectItems = append(subSelectItems, node.SelectStmt.HavingClause)
				}

				switch node.SelectStmt.Op {
				case SetOperation_SETOP_NONE:
					if len(node.SelectStmt.FromClause) > 0 {
						// From sub selects
						for _, e := range node.SelectStmt.FromClause {
							fromClause = append(fromClause, TaggedNode{Item: e.GetNode(), T: TagSelect})
						}
					}
				case SetOperation_SETOP_UNION, SetOperation_SETOP_INTERSECT, SetOperation_SETOP_EXCEPT:
					if node.SelectStmt.Larg != nil {
						larg := node.SelectStmt.Larg
						stms = append(stms, &Node{Node: &Node_SelectStmt{SelectStmt: larg}})
					}
					if node.SelectStmt.Rarg != nil {
						rarg := node.SelectStmt.Rarg
						stms = append(stms, &Node{Node: &Node_SelectStmt{SelectStmt: rarg}})
					}
				}
				if node.SelectStmt.WithClause != nil {
					cteStms, _cteNames := statements_and_cte_names_for_with_clause(node.SelectStmt.WithClause)
					for _, e := range _cteNames {
						tpr.CteNames[e] = true
					}
					stms = append(stms, cteStms...)
				}
			case *Node_InsertStmt:
				// TODO: value = statement.public_send(statement.node)
				fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.InsertStmt.Relation}, T: TagDML})
				if node.InsertStmt.SelectStmt != nil {
					stms = append(stms, node.InsertStmt.SelectStmt)
				}
				if node.InsertStmt.WithClause != nil {
					cteStms, _cteNames := statements_and_cte_names_for_with_clause(node.InsertStmt.WithClause)
					for _, e := range _cteNames {
						tpr.CteNames[e] = true
					}
					stms = append(stms, cteStms...)
				}
			case *Node_UpdateStmt:
				// TODO: value = statement.public_send(statement.node)
				fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.UpdateStmt.Relation}, T: TagDML})

				for _, e := range node.UpdateStmt.FromClause {
					fromClause = append(fromClause, TaggedNode{Item: e.GetNode(), T: TagSelect})
				}
				subSelectItems = append(subSelectItems, node.UpdateStmt.TargetList...)

				if node.UpdateStmt.WhereClause != nil {
					subSelectItems = append(subSelectItems, node.UpdateStmt.WhereClause)
				}
				if node.UpdateStmt.WithClause != nil {
					cteStms, _cteNames := statements_and_cte_names_for_with_clause(node.UpdateStmt.WithClause)
					for _, e := range _cteNames {
						tpr.CteNames[e] = true
					}
					stms = append(stms, cteStms...)
				}
			case *Node_DeleteStmt:
				// TODO: value = statement.public_send(statement.node)
				fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.DeleteStmt.Relation}, T: TagDML})
				if node.DeleteStmt.WhereClause != nil {
					subSelectItems = append(subSelectItems, node.DeleteStmt.WhereClause)
				}
				for _, e := range node.DeleteStmt.UsingClause {
					fromClause = append(fromClause, TaggedNode{Item: e.GetNode(), T: TagSelect})
				}
				if node.DeleteStmt.WithClause != nil {
					cteStms, _cteNames := statements_and_cte_names_for_with_clause(node.DeleteStmt.WithClause)
					for _, e := range _cteNames {
						tpr.CteNames[e] = true
					}
					stms = append(stms, cteStms...)
				}
			case *Node_CopyStmt:
				if node.CopyStmt.Relation != nil {
					fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.CopyStmt.Relation}, T: TagDML})
				}
				stms = append(stms, node.CopyStmt.Query)
			case *Node_AlterTableStmt:
				fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.AlterTableStmt.Relation}, T: TagDDL})
			case *Node_CreateStmt:
				fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.CreateStmt.Relation}, T: TagDDL})
			case *Node_CreateTableAsStmt:
				if node.CreateTableAsStmt.Into != nil && node.CreateTableAsStmt.Into.Rel != nil {
					fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.CreateTableAsStmt.Into.Rel}, T: TagDML})
				}
				if node.CreateTableAsStmt.Query != nil {
					stms = append(stms, node.CreateTableAsStmt.Query)
				}
			case *Node_TruncateStmt:
				for _, e := range node.TruncateStmt.Relations {
					fromClause = append(fromClause, TaggedNode{Item: e.GetNode(), T: TagDDL})
				}
			case *Node_ViewStmt:
				fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.ViewStmt.View}, T: TagDDL})
				stms = append(stms, node.ViewStmt.Query)
			case *Node_IndexStmt:
				fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.IndexStmt.Relation}, T: TagDDL})
				for _, e := range node.IndexStmt.IndexParams {
					if expr := e.GetIndexElem().Expr; expr != nil {
						subSelectItems = append(subSelectItems, e.GetIndexElem().Expr)
					}
				}
				if node.IndexStmt.WhereClause != nil {
					subSelectItems = append(subSelectItems, node.IndexStmt.WhereClause)
				}
			case *Node_CreateTrigStmt:
				fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.CreateTrigStmt.Relation}, T: TagDDL})
			case *Node_RuleStmt:
				fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.RuleStmt.Relation}, T: TagDDL})
			case *Node_VacuumStmt:
				for _, e := range node.VacuumStmt.Rels {
					if n := e.GetVacuumRelation(); n != nil {
						fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: n.Relation}, T: TagDDL})
					}
				}
			case *Node_RefreshMatViewStmt:
				fromClause = append(fromClause, TaggedNode{Item: &Node_RangeVar{RangeVar: node.RefreshMatViewStmt.Relation}, T: TagDDL})
			case *Node_DropStmt:
				var objects [][]string
				for _, e := range node.DropStmt.Objects {
					buf := []string{}
					switch n := e.GetNode().(type) {
					case *Node_List:
						for _, e := range n.List.Items {
							if e.GetString_() != nil {
								buf = append(buf, e.GetString_().GetSval())
							}
						}
					case *Node_String_:
						buf = append(buf, n.String_.GetSval())
					}
					objects = append(objects, buf)
				}
				switch node.DropStmt.RemoveType {
				case ObjectType_OBJECT_TABLE:
					for _, e := range objects {
						tpr.tables = append(tpr.tables, TaggedNode{Name: strings.Join(e, "."), T: TagDDL})
					}
				case ObjectType_OBJECT_RULE, ObjectType_OBJECT_TRIGGER:
					for _, e := range objects {
						if len(e) > 1 {
							e = e[:len(e)-1]
						}
						tpr.tables = append(tpr.tables, TaggedNode{Name: strings.Join(e, "."), T: TagDDL})
					}
				case ObjectType_OBJECT_FUNCTION:
					obj := node.DropStmt.Objects[0].GetObjectWithArgs()
					var name string
					for _, e := range obj.Objname {
						if s := e.GetString_(); s != nil {
							name += s.GetSval() + "."
						}
					}
					name = name[:len(name)-1]
					tpr.functions = append(tpr.functions, TaggedNode{Name: name, T: TagDDL})
				}
			case *Node_GrantStmt:
				switch node.GrantStmt.Objtype {
				case ObjectType_OBJECT_COLUMN:
					// TODO FIXME Not implemented in Ruby
				case ObjectType_OBJECT_SEQUENCE:
				// TODO FIXME Not implemented in Ruby
				case ObjectType_OBJECT_TABLE:
					for _, e := range node.GrantStmt.Objects {
						fromClause = append(fromClause, TaggedNode{Item: e.GetNode(), T: TagDDL})
					}
				}
			case *Node_LockStmt:
				for _, e := range node.LockStmt.Relations {
					fromClause = append(fromClause, TaggedNode{Item: e.GetNode(), T: TagDDL})
				}
			case *Node_ExplainStmt:
				stms = append(stms, node.ExplainStmt.Query)
			case *Node_CreateFunctionStmt:
				var funcNames string
				for _, e := range node.CreateFunctionStmt.Funcname {
					if n := e.GetString_(); n != nil {
						funcNames = funcNames + n.GetSval() + "."
					}
				}
				funcNames = funcNames[:len(funcNames)-1]
				tpr.functions = append(tpr.functions, TaggedNode{Name: funcNames, T: TagDDL})
			case *Node_RenameStmt:
				if node.RenameStmt.RenameType == ObjectType_OBJECT_FUNCTION {
					var origName string
					for _, e := range node.RenameStmt.Object.GetObjectWithArgs().Objname {
						if n := e.GetString_(); n != nil {
							origName = origName + n.GetSval() + "."
						}
					}
					origName = origName[:len(origName)-1]
					tpr.functions = append(tpr.functions, TaggedNode{Name: origName, T: TagDDL},
						TaggedNode{Name: node.RenameStmt.Newname, T: TagDDL})
				}

			case *Node_PrepareStmt:
				stms = append(stms, node.PrepareStmt.Query)

			default:
				// Extract typename of the node and pretty print it for easy debugging
				if debugLog && e != nil {
					t := reflect.TypeOf(e).Elem()
					log.Printf("Unhandled node type: %s", t.Name())
				}
			}
		}
		if len(subSelectItems) > 0 {
			e := subSelectItems[0]
			subSelectItems = subSelectItems[1:]
			if e == nil {
				continue
			}
			switch n := e.GetNode().(type) {
			case *Node_List:
				subSelectItems = append(subSelectItems, n.List.Items...)
			case *Node_AExpr:
				if n.AExpr.Lexpr != nil {
					subSelectItems = append(subSelectItems, n.AExpr.Lexpr)
				}
				if n.AExpr.Rexpr != nil {
					subSelectItems = append(subSelectItems, n.AExpr.Rexpr)
				}
			case *Node_BoolExpr:
				subSelectItems = append(subSelectItems, n.BoolExpr.Args...)
			case *Node_CoalesceExpr:
				subSelectItems = append(subSelectItems, n.CoalesceExpr.Args...)
			case *Node_MinMaxExpr:
				subSelectItems = append(subSelectItems, n.MinMaxExpr.Args...)
			case *Node_ResTarget:
				subSelectItems = append(subSelectItems, n.ResTarget.Val)
			case *Node_SubLink:
				stms = append(stms, n.SubLink.Subselect)
			case *Node_FuncCall:
				var funcNames string
				for _, e := range n.FuncCall.Funcname {
					if n := e.GetString_(); n != nil {
						funcNames = funcNames + n.GetSval() + "."
					}
				}
				funcNames = funcNames[:len(funcNames)-1]
				for _, _node := range n.FuncCall.Args {
					switch node := _node.GetNode().(type) {
					case *Node_ColumnRef:
						var strbuf []string
						for _, e := range node.ColumnRef.Fields {
							switch node := e.GetNode().(type) {
							case *Node_String_:
								strbuf = append(strbuf, node.String_.GetSval())
							case *Node_AStar:
								strbuf = append(strbuf, node.AStar.String())
							}
						}
						var argname string
						if len(strbuf) == 1 {
							argname = strbuf[0]
						} else if len(strbuf) == 2 {
							if a, ok := tpr.Aliases[strbuf[0]]; ok {
								argname = strings.Join([]string{a, strbuf[1]}, ".")
							} else {
								argname = strings.Join([]string{strbuf[0], strbuf[1]}, ".")
							}
						} else {
							log.Printf("ColumnRef doesn't have 1-2 fields, got %d", len(strbuf))
							return
						}
						tpr.functionsArgs = append(tpr.functions, TaggedNode{Item: node, Name: argname, T: TagFuncArg})
					default:

						subSelectItems = append(subSelectItems, _node)
					}
				}
				tpr.functions = append(tpr.functions, TaggedNode{Name: funcNames, T: TagCall})
			case *Node_CaseExpr:
				for _, e := range n.CaseExpr.Args {
					subSelectItems = append(subSelectItems, e.GetCaseWhen().Expr)
				}
				for _, e := range n.CaseExpr.Args {
					subSelectItems = append(subSelectItems, e.GetCaseWhen().Result)
				}
				subSelectItems = append(subSelectItems, n.CaseExpr.Defresult)
			case *Node_TypeCast:
				subSelectItems = append(subSelectItems, n.TypeCast.Arg)
			default:
				// Extract typename of the node and pretty print it for easy debugging
				if debugLog && e != nil {
					t := reflect.TypeOf(e).Elem()
					log.Printf("Unhandled node type: %s", t.Name())
				}
			}
		}
		if len(fromClause) > 0 {
			e := fromClause[0]
			fromClause = fromClause[1:]
			switch n := e.Item.(type) {
			case *Node_JoinExpr:
				if n.JoinExpr.Larg != nil {
					larg := n.JoinExpr.Larg
					fromClause = append(fromClause, TaggedNode{Item: larg.GetNode(), T: e.T})
				}
				if n.JoinExpr.Rarg != nil {
					rarg := n.JoinExpr.Rarg
					fromClause = append(fromClause, TaggedNode{Item: rarg.GetNode(), T: e.T})
				}
				subSelectItems = append(subSelectItems, n.JoinExpr.Quals)
			case *Node_RowExpr:
				for _, m := range n.RowExpr.Args {
					fromClause = append(fromClause, TaggedNode{Item: m.GetNode(), T: e.T})
				}
			case *Node_RangeVar:
				if n.RangeVar.Schemaname == "" && tpr.CteNames[n.RangeVar.Relname] {
					continue
				}
				var table string
				if n.RangeVar.Schemaname != "" {
					if n.RangeVar.Relname != "" {
						table = n.RangeVar.Schemaname + "." + n.RangeVar.Relname
					} else {
						table = n.RangeVar.Schemaname
					}
				} else {
					table = n.RangeVar.Relname
				}
				tpr.tables = append(tpr.tables, TaggedNode{Item: &Node_RangeVar{RangeVar: &RangeVar{
					Location:       n.RangeVar.Location,
					Schemaname:     n.RangeVar.Schemaname,
					Relname:        n.RangeVar.Relname,
					Inh:            n.RangeVar.Inh,
					Relpersistence: n.RangeVar.Relpersistence,
				}}, Name: table, T: e.T})
				if n.RangeVar.Alias != nil {
					tpr.Aliases[n.RangeVar.Alias.Aliasname] = table
				}
			case *Node_RangeSubselect:
				stms = append(stms, n.RangeSubselect.Subquery)
			case *Node_RangeFunction:
				subSelectItems = append(subSelectItems, n.RangeFunction.Functions...)
			default:
				// Extract typename of the node and pretty print it for easy debugging
				if debugLog {
					t := reflect.TypeOf(e).Elem()
					log.Printf("Unhandled node type: %s", t.Name())
				}
			}
		}
	}

	return
}
