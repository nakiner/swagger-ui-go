package swaggerui

import (
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

// HTTPHandler swagger-ui handler.
func HTTPHandler() http.Handler {
	return http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo})
}
