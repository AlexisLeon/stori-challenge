#!/bin/bash

# Get the number of commits
num_commits=$(git rev-list --count HEAD)

# Loop through each commit
for (( i=1; i<=$num_commits; i++ ))
do
    # Get the commit hash
    commit_hash=$(git rev-list --reverse HEAD | sed -n ${i}p)

    days_offset=108
    # offset first 7 commits
    if test $i -lt 8; then
        days_offset=111
    fi

    # Calculate the new date
    current_date=$(git show -s --format=%cd --date=iso $commit_hash)
    new_date=$(date -v +"$days_offset"d -j -f "%Y-%m-%d %H:%M:%S %z" "$current_date" "+%Y-%m-%d %H:%M:%S %z")

    # Update the author date and commit date
    GIT_COMMITTER_DATE="$new_date" git commit --amend --no-edit --date "$new_date" $commit_hash
    # echo "commit $i: +$days_offset days - $current_date -> $new_date"
done
