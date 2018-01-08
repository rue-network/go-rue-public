#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
ruedir="$workspace/src/github.com/Rue-Foundation"
if [ ! -L "$ruedir/go-rue" ]; then
    mkdir -p "$ruedir"
    cd "$ruedir"
    ln -s ../../../../../. go-rue
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$ruedir/go-rue"
PWD="$ruedir/go-rue"

# Launch the arguments with the configured environment.
exec "$@"
