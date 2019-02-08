#!/usr/bin/env bash

find workdir/crashers -type f ! -name "*.*" -print | while IFS= read -r FILE; do
	PANIC="$(go run run.go < $FILE 2>&1)"
	if [ $? != 0 ]; then
		echo "Still crashing with input file $FILE: $PANIC"
	fi
done
