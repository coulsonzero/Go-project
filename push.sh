#! /bin/bash
set -e

git add . & git status
# shellcheck disable=SC2162
read -t 60 -p "[master] Enter commit >>> " message
if [ -n "$message" ]; then
  git commit -m "[commit]: $message"
  git pull origin master
  git push origin master
  exit 0
else
  git reset
  echo "[master] You haven't entered any comments !"
  exit 1
fi
