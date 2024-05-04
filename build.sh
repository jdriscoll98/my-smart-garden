#!/bin/bash

echo "Building client..."
cd client
go build -o clientApp
cd ..

echo "Building server..."
cd server
go build -o serverApp
cd ..

echo "Build complete."
