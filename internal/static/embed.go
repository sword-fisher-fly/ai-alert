package static

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist/*
var distFS embed.FS

func GetFileSystem() http.FileSystem {
	subFS, err := fs.Sub(distFS, "dist")
	if err != nil {
		panic(err)
	}
	return http.FS(subFS)
}

func GetStaticFileSystem() http.FileSystem {
	subFS, err := fs.Sub(distFS, "dist/static")
	if err != nil {
		panic(err)
	}
	return http.FS(subFS)
}

func GetFS() fs.FS {
	subFS, err := fs.Sub(distFS, "dist")
	if err != nil {
		panic(err)
	}
	return subFS
}

func ReadFile(name string) ([]byte, error) {
	return distFS.ReadFile("dist/" + name)
}
