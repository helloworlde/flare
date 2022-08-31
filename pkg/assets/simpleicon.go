package assets

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/flare/internal/logger"
	"github.com/soulteary/memfs"
)

var MemFs *memfs.FS

const _ASSETS_ICON_DIR = "assets/icons"
const _ASSETS_ICON_URI = "/" + _ASSETS_ICON_DIR

//go:embed icons
var SimpleIconsAssets embed.FS
var simpleIcons map[string]string
var log = logger.GetLogger()

func InitIcon() {
	MemFs = memfs.New()
	err := MemFs.MkdirAll(_ASSETS_ICON_DIR, 0777)

	if err != nil {
		panic(err)
	}

	simpleIcons = make(map[string]string)

	dirs, err := SimpleIconsAssets.ReadDir("icons")
	if err != nil {
		panic(err)
	}
	for i := range dirs {
		file := dirs[i]
		simpleIcons[strings.ToLower(file.Name())] = file.Name()
	}
	log.Println("初始化 Simple icon 共 ", len(simpleIcons), " 个")
}

func RegisterIconRouting(router *gin.Engine) {
	mdiExample, _ := fs.Sub(SimpleIconsAssets, "icons")
	router.StaticFS(_ASSETS_ICON_URI, http.FS(mdiExample))
}

const _EMPTY_ICON = ""

func GetIconByName(name string) string {
	if name == "" {
		return _EMPTY_ICON
	}
	svgFile := filepath.Join(_ASSETS_ICON_DIR, strings.ToLower(name)+".svg")
	if _, ok := simpleIcons[strings.ToLower(name)+".svg"]; ok {
		return `<img src="/` + svgFile + `" width="68" height="68" alt="">`
	} else {
		log.Println("Simple icon '" + name + "' 不存在")
		return _EMPTY_ICON
	}
}
