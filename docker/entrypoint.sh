#!/bin/bash
set -e

alias tmux="tmux -2"

SESSION="tvtio"

tmux new-session -d -s $SESSION
tmux rename-window "vim"
tmux send-keys "nvim" C-m
tmux new-window -t $SESSION:1 -n "shell"
tmux select-window -t $SESSION:0
tmux attach-session -t $SESSION
