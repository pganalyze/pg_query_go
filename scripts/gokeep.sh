#!/bin/bash -e

c_dir="parser/include"

go_module="$(go list -m)"
go_imports=""

while IFS='' read -r c_dir; do
    go_imports+="	_ \"${go_module}/${c_dir}\""$'\n'
    cat <<EOF >"${c_dir}/gokeep.go"
//go:build required
// +build required

// package gokeep prevents go tooling from stripping the C dependencies.
package gokeep
EOF
done < <(find ${c_dir} -type f \( -name '*.c' -o -name '*.h' \) -exec dirname {} \; | sort -u)

cat <<EOF >"parser/build_cgo.go"
//go:build required
// +build required

package parser

// This file exists purely to prevent the Golang toolchain from stripping
// away the C source directories and files when \`go mod vendor\` is used
// to populate a \`vendor/\` directory of a project depending on \`${go_module}\`.
//
// How it works:
//  - Every directory which only includes C source files receives a gokeep.go file.
//  - Every directory we want to preserve is included here as a _ import.
//  - This file is given a build tag to exclude it from the regular build.
import (
	// Prevent Go tooling from stripping out the C source files.
${go_imports})
EOF
