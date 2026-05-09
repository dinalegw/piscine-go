#!/bin/bash

echo "🔍 Searching for the mystery folder..."

MYSTERY_PATH=$(find ~ -type d -name "mystery" 2>/dev/null | head -n 1)

if [ -z "$MYSTERY_PATH" ]; then
    echo "❌ No mystery folder found on your system."
    echo "Make sure the exercise repository is cloned correctly."
    exit 1
fi

echo "✅ Mystery folder found at:"
echo "$MYSTERY_PATH"
echo

cd "$MYSTERY_PATH" || { echo "❌ Failed to enter directory."; exit 1; }

echo "📁 You are now inside:"
pwd
echo

echo "📂 Listing folder contents:"
echo "-------------------------------------------"
ls -R
echo "-------------------------------------------"

echo
echo "🎉 Done! Copy the above output and send it to ChatGPT so it can build your teacher.sh"
echo
