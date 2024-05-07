#! /bin/bash
find . -type f -name "file*.sh" | tr -d '.sh' | tr -d '/' | sort -r 