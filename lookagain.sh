#! /bin/bash
find . -name '*.sh' | sed 's/.sh//' | sed 's/.*\///'| sort -r