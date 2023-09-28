# swagger-ui-go

#### Provides http handler for swagger ui static files

Generate new swagger:
1. Place new .js, .png, .html, .css files inside the folder
2. `brew install go-bindata` or use newest https://github.com/kevinburke/go-bindata
3. `go-bindata -o bindata.go *.css *.js *.png *.html`
4. inside the bindata.go set package to `swaggerui`

How to use:
```go
mux.Mount("/swagger", http.StripPrefix("/swagger", swagger.HTTPHandler()))
```