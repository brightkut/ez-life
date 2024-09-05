#!/bin/bash
echo "Installing Hyper..."

# Check if user is root
if [ "$EUID" -ne 0 ]; then
  echo "Please run as root"
  exit
fi

brew update
brew install --cask hyper

echo "Installation complete. You can run Hyper using terminal."