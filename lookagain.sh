#! /bin/bash
find . -type f -name "*.sh" | tr -d '.sh' | tr -d '/' | sort -r 