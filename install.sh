#!/bin/bash

# Build the Go program
go build pingpong.go

# Move the binary to /usr/local/bin
mv pingpong /usr/local/bin/

# Install the man page
sudo mkdir -p /usr/local/share/man/man1
sudo cp pingpong.1 /usr/local/share/man/man1/

echo "Installation complete!"