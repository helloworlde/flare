package mdi

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/memfs"
)

var MemFs *memfs.FS

const _ASSETS_ICON_DIR = "simpleicon/icons"
const _ASSETS_ICON_URI = "/" + _ASSETS_ICON_DIR

//go:embed icons
var SimpleIconsAssets embed.FS

func Init() {
	MemFs = memfs.New()
	err := MemFs.MkdirAll(_ASSETS_ICON_DIR, 0777)

	if err != nil {
		panic(err)
	}
}

func RegisterRouting(router *gin.Engine) {
	mdiExample, _ := fs.Sub(SimpleIconsAssets, "icons")
	router.StaticFS(_ASSETS_ICON_URI, http.FS(mdiExample))
}

const _EMPTY_ICON = ""

func GetIconByName(name string) string {
	if name == "" {
		return _EMPTY_ICON
	}
	svgFile := filepath.Join(_ASSETS_ICON_DIR, name+".svg")

	return `<img src="/` + svgFile + `" width="68" height="68" alt="">`
}
