package fspkg

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/wmem/go-cmlib/strpkg"
)

type FileFilter struct {
	//  dir
	//	*.xxx
	//	xxx/.../yyy/*.zzz
	//  dir/file
	Dirs   strpkg.StringSet
	Files  strpkg.StringSet
	Suffix strpkg.StringSet
}

// PathFilesCurHasSuffix 当前目录下带指定后缀的文件
func PathFilesCurHasSuffix(pathName string, suffixes strpkg.StringSet) strpkg.StringSet {

	files := strpkg.NewStringSet()
	pathName = path.Clean(pathName)

	infos, err := os.ReadDir(pathName)
	if err != nil {
		return files
	}

	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		if !suffixes.IsEmpty() && !suffixes.IsExist(path.Ext(info.Name())) {
			continue
		}
		files.Add(pathName + "/" + info.Name())
	}
	return files
}

// PathFilesCur 获取当前目录下的文件
func PathFilesCur(pathName string) strpkg.StringSet {
	return PathFilesCurHasSuffix(pathName, strpkg.NewStringSet())
}

// PathFilesALLHasSuffix 递归目录下带指定后缀的文件
func PathFilesALLHasSuffix(pathName string, suffixes strpkg.StringSet) strpkg.StringSet {

	pathName = path.Clean(pathName)
	files := strpkg.NewStringSet()

	if _, err := os.Stat(pathName); err != nil {
		return files
	}
	filepath.Walk(pathName, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !suffixes.IsExist(path.Ext(filePath)) {
			return nil
		}

		files.Add(filePath)
		return nil
	})

	return files
}

// 递归获取目录下所有文件
func pathFilesALl(pathName string) strpkg.StringSet {
	return PathFilesALLHasSuffix(pathName, strpkg.NewStringSet())
}

func NewFileFilter() FileFilter {
	return FileFilter{strpkg.NewStringSet(), strpkg.NewStringSet(), strpkg.NewStringSet()}
}

func FileFilterInit(pathRoot string, infos strpkg.StringSet) FileFilter {

	filter := NewFileFilter()
	pathRoot = path.Clean(pathRoot)

	infos.ForEach(func(info string) {
		// *.xxx
		if strings.HasPrefix(info, "*.") {
			filter.Suffix.Add(strings.Replace(info, "*.", ".", 1))

			// xxx/yyy/*sss.c
		} else if strpkg.Match("*/?*[^/]", info) {
			dir := path.Clean(pathRoot + "/" + path.Dir(info))
			if info, err := os.Stat(dir); err != nil || !info.IsDir() {
				return
			}
			if fs, err := os.ReadDir(dir); err == nil {
				for _, file := range fs {
					if file.IsDir() {
						return
					}
					if strpkg.Match(pathRoot+"/"+info, dir+"/"+file.Name()) {
						filter.Files.Add(dir + "/" + file.Name())
					}
				}
			}

			//dir or file
		} else {
			filePath := path.Clean(pathRoot + "/" + info)
			if info, err := os.Stat(filePath); err == nil {
				if info.IsDir() {
					filter.Dirs.Add(filePath)
				} else {
					filter.Files.Add(filePath)
				}
			}
		}
	})

	return filter
}

// SuffixFiles 获取指定后缀代表的文件
func (f *FileFilter) SuffixFiles(rootPath string) strpkg.StringSet {
	return PathFilesALLHasSuffix(rootPath, f.Suffix)
}
