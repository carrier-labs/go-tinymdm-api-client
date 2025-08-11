#!/bin/bash

set -e

# Ensure on main branch
branch=$(git rev-parse --abbrev-ref HEAD)
if [ "$branch" != "main" ]; then
  echo "You must be on the main branch to tag a release."
  exit 1
fi

# Get latest tag (or default to v0.0.0)
latest=$(git tag --list 'v*' --sort=-v:refname | head -n1)
if [ -z "$latest" ]; then
  latest="v0.0.0"
fi
echo "Latest tag: $latest"

# Parse version
ver=${latest#v}
IFS='.' read -r major minor patch <<< "$ver"

echo "What do you want to increment? (major/minor/patch)"
read -r part

case "$part" in
  major)
    major=$((major + 1))
    minor=0
    patch=0
    ;;
  minor)
    minor=$((minor + 1))
    patch=0
    ;;
  patch)
    patch=$((patch + 1))
    ;;
  *)
    echo "Invalid choice."
    exit 1
    ;;
esac

new_tag="v$major.$minor.$patch"
echo "Tagging as $new_tag"
git tag "$new_tag"
git push origin "$new_tag"

echo "Tagging as latest"
git tag -f latest
# Use --force to update remote 'latest' tag if it already exists
git push --force origin latest