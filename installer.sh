#!/bin/bash

go build .

sudo mv tree-tool /usr/local/bin/tree

if ! grep -q "$HOME/bin" ~/.bashrc; then
    echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
    source ~/.bashrc
fi

echo "Tree installed!"
