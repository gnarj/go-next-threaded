#!/bin/zsh

# Open VSCode
open -a "Visual Studio Code" .

# Start yarn in a new terminal window
osascript -e 'tell app "Terminal" to do script "cd /Users/rj/Development/go-next-threaded/services/frontend && yarn dev"'

# Start CompileDaemon in a new terminal window
osascript -e 'tell app "Terminal" to do script "cd /Users/rj/Development/go-next-threaded/services/backend && /Users/rj/go/bin/CompileDaemon -command=\"./backend\""'
