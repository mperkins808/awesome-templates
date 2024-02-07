# react-go

A Fully set up template to create a React App that is Served by a Go Server

| **Component** | **Tooling**                      | **Hot reloading**                |
| ------------- | -------------------------------- | -------------------------------- |
| Frontend      | React. Bundled by Vite           | Yes                              |
| Backend       | A Go-chi Server                  | Yes                              |
| CSS           | modules                          | No. Requires refreshing the page |
| Page routing  | react-router                     | N/A                              |
| API endpoints | Yes. Check `/server/cmd/main.go` | N/A                              |

## How to Use

### Setup

Installing dependencies

```bash
make dep
```

### For Development

Starting the server

```bash
make server
```

Starting the vite dev server

```bash
make app
```

### For Production

Building the Docker image

```bash
make build
```

Running the Docker image

```bash
make run
```
