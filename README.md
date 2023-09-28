# swagger-ui-go

#### Provides http handler for swagger ui static files

Generate new swagger:
1. Place new .js, .png, .html, .css files inside the folder
2. `brew install go-bindata` or use newest https://github.com/kevinburke/go-bindata
3. `go-bindata -pkg swaggerui -o bindata.go *.css *.js *.png *.html`

How to use:
```go
mux.Mount("/swagger", http.StripPrefix("/swagger", swagger.HTTPHandler()))
```