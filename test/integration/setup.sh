#!/usr/bin/env bash
PROJECT_ROOT=${TESTDIR}/../..

# build flextime binary
cd "$ROJECT_ROOT" || exit 1
go build
cd "$TESTDIR" || exit 1

export FLEXTIME="$TESTDIR"/../../flextime