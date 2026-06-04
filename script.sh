#!/bin/bash

echo "Compilando executable de go"
echo "Compilando..."
go build -ldflags="-s -w" -o "build/tasks-cli"
echo "El ejecutable se encuentra en la carpeta /build"
