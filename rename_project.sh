#!/bin/bash

# Check if the correct number of arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: ./rename_project.sh <old-module-name> <new-module-name>"
    exit 1
fi

OLD_MODULE_NAME=$1
NEW_MODULE_NAME=$2
OLD_DIRECTORY_NAME=$(basename "$OLD_MODULE_NAME")
NEW_DIRECTORY_NAME=$(basename "$NEW_MODULE_NAME")

# Rename module in go.mod
sed -i '' "s|$OLD_MODULE_NAME|$NEW_MODULE_NAME|g" go.mod

# Replace all occurrences of the old module name with the new one in the codebase
grep -rl "$OLD_MODULE_NAME" . | xargs sed -i '' "s|$OLD_MODULE_NAME|$NEW_MODULE_NAME|g"

# Rename the root directory of the project
cd ..
mv "$OLD_DIRECTORY_NAME" "$NEW_DIRECTORY_NAME"
cd "$NEW_DIRECTORY_NAME"

echo "Project name and module paths have been updated from $OLD_MODULE_NAME to $NEW_MODULE_NAME"
echo "Project directory has been renamed from $OLD_DIRECTORY_NAME to $NEW_DIRECTORY_NAME"