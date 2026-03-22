#!/bin/bash

# Folders
EXERCISES="exercises"
EXAMPLES="examples"
NOTES="NOTES.md"

# Header
cat <<EOL > $NOTES
# LearnGo Notes

This file contains my complete Golang practice & experiment notes.

## 📁 Repo Structure
- exercises/ – function exercises, loops, calculations
- examples/ – Go code examples, file I/O, HTTP, CLI tools
- data/ – data / txt files

---

## 📊 Repo Stats
EOL

# Count files
EX_NUM=$(ls $EXERCISES/*.go 2>/dev/null | wc -l)
EXAM_NUM=$(ls $EXAMPLES/*.go 2>/dev/null | wc -l)
TOTAL=$((EX_NUM + EXAM_NUM))

# Write table
cat <<EOL >> $NOTES
| Category   | Number of Files | Description |
|------------|----------------|-------------|
| Exercises  | $EX_NUM        | Function, loop, calculation exercises |
| Examples   | $EXAM_NUM      | File I/O, HTTP, CLI tools examples |
| Total      | $TOTAL         | Total Go files in the repo |
EOL

echo "NOTES.md updated successfully!"
