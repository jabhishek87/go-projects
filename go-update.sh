#!/usr/bin/bash

# function GO_UPD(){
#   local GO_VER=$(curl -s https://go.dev/dl/\?mode\=json | jq -r '.[0].version')
#   echo -e " + Installing GO version: ${GO_VER} linux amd64"
# #   sudo apt install -qq -y golang-go
#   wget -q "https://golang.org/dl/${GO_VER}.linux-amd64.tar.gz"
#   sudo tar xf "${GO_VER}".linux-amd64.tar.gz -C /usr/local
#   [ -f "$HOME/.bashrc" ]&& echo -e "Info: .bashrc exists!\n" || echo -e "Error: .bashrc doesn't exist! Creating... $(> $HOME/.bashrc)\n"
#   echo -e '\nexport GOROOT=/usr/local/go\nexport GOPATH=$HOME/go\nexport PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> $HOME/.bashrc
#   source "$HOME"/.bashrc
# }
function GO_UPD(){
    local GO_VER=$(curl -s https://go.dev/dl/\?mode\=json | jq -r '.[0].version')
    echo -e " + Installing GO version: ${GO_VER} linux amd64"
    wget -q "https://golang.org/dl/${GO_VER}.linux-amd64.tar.gz"
    sudo tar xf "${GO_VER}".linux-amd64.tar.gz -C /usr/local
}
GO_UPD