# React + TypeScript + Vite + Golang

> WORK IN PROGRESS

This template provides a minimal setup for application consisting of backend writen in Go and front-end written with React.

In dev mode Golang server acts as a proxy for Vite server
In prod Golang server hosts entire application

## How to run

Go server compiles to single exec file. To run, compile front end project and then exec compiled server.

```
npm run build
go build -o ./bin/server
cp -rf ./dist ./bin
```
