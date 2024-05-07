#! /bin/bash
find . -name "file*.sh" | tr .h/s 0000 | sed 's/0/ /g' | tr -d ' ' | sort -r