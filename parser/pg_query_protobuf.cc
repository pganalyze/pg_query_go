#include <stdio.h>
#include <stdlib.h>

#include <iostream>
#include <fstream>
#include <string>
#include <protobuf/parse_tree.pb.h>
#include <google/protobuf/util/json_util.h>

#include "pg_query_protobuf.h"

extern "C"
{
#include "postgres.h"
#include <ctype.h>
#include "access/relation.h"
#include "nodes/parsenodes.h"
#include "nodes/plannodes.h"
#include "nodes/value.h"
#include "utils/datum.h"
}

#define OUT_NODE(typename, typename_cast, fldname) \
  { \
    pg_query::typename *fldname = new pg_query::typename(); \
    out->set_allocated_##fldname(fldname); \
    _out##typename(fldname, (const typename_cast *) obj); \
  }

#define WRITE_INT_FIELD(outname, fldname) \
	if (node->fldname != 0) { \
    out_node->set_##outname(node->fldname); \
	}

#define WRITE_UINT_FIELD(outname, fldname) \
	if (node->fldname != 0) { \
    out_node->set_##outname(node->fldname); \
	}

#define WRITE_LONG_FIELD(outname, fldname) \
	if (node->fldname != 0) { \
    out_node->set_##outname(node->fldname); \
	}

#define WRITE_CHAR_FIELD(outname, fldname) \
	if (node->fldname != 0) { \
    out_node->set_##outname({node->fldname}); \
	}

#define WRITE_ENUM_FIELD(typename, outname, fldname) \
	out_node->set_##outname((pg_query::typename) node->fldname); \

#define WRITE_FLOAT_FIELD(outname, fldname) \
	out_node->set_##outname(node->fldname); \

#define WRITE_BOOL_FIELD(outname, fldname) \
	out_node->set_##outname(node->fldname); \

#define WRITE_STRING_FIELD(outname, fldname) \
	if (node->fldname != NULL) { \
	  out_node->set_##outname(node->fldname); \
	}

#define WRITE_LIST_FIELD(outname, fldname) \
	if (node->fldname != NULL) { \
    const ListCell *lc; \
    foreach(lc, node->fldname) \
    { \
      _outNode(out_node->add_##outname(), lfirst(lc)); \
    } \
  }

#define WRITE_BITMAPSET_FIELD(outname, fldname) // FIXME

#define WRITE_NODE_FIELD(outname, fldname) \
	if (true) { \
		  out_node->set_allocated_##fldname(new pg_query::Node()); \
      _outNode(out_node->mutable_##outname(), &node->fldname); \
  	}

#define WRITE_NODE_PTR_FIELD(outname, fldname) \
	if (node->fldname != NULL) { \
    out_node->set_allocated_##outname(new pg_query::Node()); \
    _outNode(out_node->mutable_##outname(), node->fldname); \
	}

#define WRITE_SPECIFIC_NODE_FIELD(typename, outname, fldname) \
  out_node->set_allocated_##outname(new pg_query::typename()); \
  _out##typename(out_node->mutable_##outname(), &node->fldname); \

#define WRITE_SPECIFIC_NODE_PTR_FIELD(typename, outname, fldname) \
  if (node->fldname != NULL) { \
    out_node->set_allocated_##outname(new pg_query::typename()); \
    _out##typename(out_node->mutable_##outname(), node->fldname); \
	}

static void _outNode(pg_query::Node*, const void *);

extern "C" char *
pg_query_nodes_to_protobuf(const void *obj)
{
  if (obj == NULL)
    return pstrdup("");

	pg_query::Node root_node;

  _outNode(&root_node, obj);

  std::string output;
  root_node.SerializeToString(&output);
  return pstrdup(output.c_str());
}

extern "C" char *
pg_query_nodes_to_protobuf_json(const void *obj)
{
  if (obj == NULL)
    return pstrdup("{}");

	pg_query::Node root_node;

  _outNode(&root_node, obj);

  std::string output;
  google::protobuf::util::MessageToJsonString(root_node.list(), &output);
  // Drop leading {"items": and trailing }
  return pstrdup(output.substr(9, output.size() - 9 - 1).c_str());
}

#include "pg_query_protobuf_defs.cc"

static void
_outList(pg_query::List* out_node, const List *node)
{
	const ListCell *lc;

	foreach(lc, node)
	{
		_outNode(out_node->add_items(), lfirst(lc));
	}
}

static void
_outIntList(pg_query::IntList* out_node, const List *node)
{
	const ListCell *lc;

	foreach(lc, node)
	{
		_outNode(out_node->add_items(), lfirst(lc));
	}
}

static void
_outOidList(pg_query::OidList* out_node, const List *node)
{
	const ListCell *lc;

	foreach(lc, node)
	{
		_outNode(out_node->add_items(), lfirst(lc));
	}
}

// TODO: Add Bitmapset

static void
_outInteger(pg_query::Integer* out_node, const Value *node)
{
  out_node->set_ival(node->val.ival);
}

static void
_outFloat(pg_query::Float* out_node, const Value *node)
{
  out_node->set_str(node->val.str);
}

static void
_outString(pg_query::String* out_node, const Value *node)
{
  out_node->set_str(node->val.str);
}

static void
_outBitString(pg_query::BitString* out_node, const Value *node)
{
  out_node->set_str(node->val.str);
}

static void
_outNull(pg_query::Null* out_node, const Value *node)
{
  // Null has no fields
}

static void
_outNode(pg_query::Node* out, const void *obj)
{
	if (obj == NULL)
	{
		// TODO
	}
	else if (IsA(obj, List))
	{
    OUT_NODE(List, List, list);
	}
	else
	{
		switch (nodeTag(obj))
		{
			case T_Integer:
        OUT_NODE(Integer, Value, integer);
				break;
			case T_Float:
        OUT_NODE(Float, Value, float_);
				break;
			case T_String:
        OUT_NODE(String, Value, string);
				break;
			case T_BitString:
        OUT_NODE(BitString, Value, bit_string);
				break;
			case T_Null:
        OUT_NODE(Null, Value, null);
				break;
			case T_IntList:
        OUT_NODE(IntList, List, int_list);
				break;
			case T_OidList:
        OUT_NODE(OidList, List, oid_list);
				break;

			#include "pg_query_protobuf_conds.cc"

			default:
        printf("could not dump unrecognized node type: %d", (int) nodeTag(obj));
				elog(WARNING, "could not dump unrecognized node type: %d",
					 (int) nodeTag(obj));

				return;
		}
	}
}

