#! /bin/bash
find . -name "file*.sh" | tr -d '.sh' | tr -d '/' | sort -r 