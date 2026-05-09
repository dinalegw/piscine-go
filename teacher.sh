#!/bin/bash

INTERVIEW_PATH=$(grep -r "blue Honda" interviews/ | head -n1 | cut -d: -f1)

INTERVIEW_NUM=$(echo "$INTERVIEW_PATH" | grep -Eo '[0-9]+')

echo "$INTERVIEW_NUM"

cat "$INTERVIEW_PATH"

echo "$MAIN_SUSPECT"
