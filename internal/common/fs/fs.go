package fs

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type File struct {
	Info *os.FileInfo
	AbsPath string
}

type FileCollection struct {
	Items []File
}

type Scanner struct {
	IncludePattern string
	ExcludePattern string
	RootDir string
	Recursive bool
}

func NewScanner(root string) *Scanner {
	scanner := &Scanner{
		RootDir:        root,
		Recursive:      true,
	}
	return scanner
}

func (s *Scanner) scanDir(coll *FileCollection, root string) {
	fmt.Printf("%s\n", root)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && s.Recursive {
			s.scanDir(coll, filepath.Join(path, string(filepath.Separator)))
			return err
		}

		//p := filepath.Join(path, info.Name())

		if len(s.ExcludePattern) > 0 {
			if exclude, _ := regexp.MatchString(s.ExcludePattern, info.Name()); exclude {
				return nil
			}
		}

		coll.Items = append(coll.Items, File{
			AbsPath: path,
			Info: &info,
		})
		return err
	})
	if err != nil {
		panic(err)
	}
}

func (s *Scanner) Run() *FileCollection {
	coll := &FileCollection{}
	err := filepath.Walk(s.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() { return nil }

		if len(s.ExcludePattern) > 0 {
			if exclude, _ := regexp.MatchString(s.ExcludePattern, info.Name()); exclude {
				return nil
			}
		}

		coll.Items = append(coll.Items, File{
			AbsPath: path,
			Info: &info,
		})
		return err
	})
	if err != nil {
		panic(err)
	}
	return coll
}
