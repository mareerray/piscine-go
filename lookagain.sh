#! /bin/bash
# find . -name "*.sh" | tr '.sh' 000 | sed 's/0/ /g' | tr -d ' ' | sort -r
find . -name "*.sh" | sed 's|.sh||' | sed 's/.*\///'| sort -r