#ifndef PG_QUERY_READFUNCS_H
#define PG_QUERY_READFUNCS_H

#include "pg_query.h"

#include "postgres.h"
#include "nodes/pg_list.h"
#include "nodes/parsenodes.h"

List * pg_query_protobuf_to_nodes(PgQueryProtobuf protobuf);
Node * pg_query_protobuf_to_node(PgQueryProtobuf protobuf);
TypeName * pg_query_protobuf_to_typename(PgQueryProtobuf protobuf);
List * pg_query_protobuf_to_list(PgQueryProtobuf protobuf);

#endif
