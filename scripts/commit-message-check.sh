#!/bin/bash

commit_regex='^(feat|fix|docs|style|refactor|test|chore): .+'

if ! grep -qE "$commit_regex" "$1"; then
    echo "❌ Invalid commit message format!"
    echo "Format: 'type: description'"
    echo "Types: feat, fix, docs, style, refactor, test, chore"
    echo "Example: 'feat: add user authentication'"
    exit 1
fi