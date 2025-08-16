# MyStack Go

This is a boilerplate for starting a Go project with the following technologies:

- Go
- Chi
- go-templ
- htmx
- Tailwind CSS
- PostgreSQL
- Docker
- Compose

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/zkurdi45/mystack-go.git
   ```

2. **IMPORTANT**: Rename the project. You need to replace all occurrences of `github.com/zkurdi45/mystack-go` with your own project name. You can use a tool like `find` and `sed` to do this:

   ```bash
   find . -type f -name '*.go' -exec sed -i 's/github.com\/zkurdi45\/mystack-go/github.com\/your-username\/your-project-name/g' {} +
   find . -type f -name '*.templ' -exec sed -i 's/github.com\/zkurdi45\/mystack-go/github.com\/your-username\/your-project-name/g' {} +
   mv go.mod go.mod.bak
   sed 's/github.com\/zkurdi45\/mystack-go/github.com\/your-username\/your-project-name/g' go.mod.bak > go.mod
   rm go.mod.bak
   ```

3. Create a `.env` file from the example:

   ```bash
   cp .env.example .env
   ```

4. Run the application:

   ```bash
   docker-compose up -d
   ```

5. The application will be available at http://localhost:8080.

## Development

This project uses `go-templ` for templating. You will need to install the `templ` CLI to generate the Go code from the `.templ` files.

### Installation

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Generating Code

To generate the Go code from the `.templ` files, run the following command:

```bash
templ generate
```

This will generate the necessary Go files in the `web/templates` directory. You should run this command whenever you make changes to the `.templ` files.
