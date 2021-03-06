package ink

import (
    "os"
    "net/http"
    "path/filepath"
)

func Static(root string) func(ctx *Context) {
    return func(ctx *Context) {
        // http.FileServer(http.Dir(root)).ink.Web
        reqURL := ctx.Req.URL.Path
        if reqURL == "" || reqURL == "/" {
            indexPath := filepath.Join(root, "index.html")
            if _, err := os.Stat(indexPath); err == nil {
                http.ServeFile(ctx.Res, ctx.Req, indexPath)
                ctx.Stop()
            }
        } else {
            fileName := root + reqURL
            filePath, _ := filepath.Abs(fileName)
            rootPath, _ := filepath.Abs(root)
            fileDir := filepath.Dir(filePath)
            f, err := os.Stat(filePath)
            if err == nil && filepath.HasPrefix(fileDir, rootPath) {
                if f.IsDir() {
                    http.ServeFile(ctx.Res, ctx.Req, filepath.Join(filePath, "index.html"))
                } else {
                    http.ServeFile(ctx.Res, ctx.Req, filePath)
                }
                ctx.Stop()
            }
        }
    }
}
