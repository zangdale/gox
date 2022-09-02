package xhttp

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path"
)

////go:embed www
// var Assets embed.FS

// FileServer web server
func FileServer(port, prefix, root, indexFile string, assets embed.FS) error {
	return http.ListenAndServe(port, AssetHandler(prefix, root, indexFile, assets))
}

var _ fs.FS = (*fsFunc)(nil)

// fsFunc 是构建 http.FileSystem 实现的简写
type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

//AssetHandler 返回一个 http.Handler ，它将从
//资产嵌入.FS。定位文件时，它将删除给定的
//请求中的前缀并将根添加到文件系统中
//查找：典型的前缀可能是 /web/，而 root 将被构建。
func AssetHandler(prefix, root, indexFile string, assets embed.FS) http.Handler {
	handler := fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)

		// 如果我们找不到资产，则返回默认的 index.html 内容
		f, err := assets.Open(assetPath)
		if os.IsNotExist(err) {
			return assets.Open(indexFile)
		}

		// 否则假设这是一个正确路由的合法请求
		return f, err
	})

	return http.StripPrefix(prefix, http.FileServer(http.FS(handler)))
}
