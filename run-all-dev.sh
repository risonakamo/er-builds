# run server and web build. need tmux already open

set -exu
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

tmux rename-window spawn

# run web build
tmux new-window -n run -c $HERE/er-builds-web
tmux send "pnpm watch" Enter

# run server
tmux split-window -h -c $HERE
tmux send "bash run-erbuilds-dev.sh" Enter