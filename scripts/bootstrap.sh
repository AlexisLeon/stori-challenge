#!/bin/bash
export HOMEBREW_NO_AUTO_UPDATE=1

brew install go postgres golang-migrate golangci-lint mockery
brew services start postgresql

echo "Done"
