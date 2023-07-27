// first-subdir-executable finds the first file executable by any user located
// in a subdirectory of the provided directory.
//
// While the error handling and resource cleanup approaches aren't scalable
// for larger applications, this program provides an example of how primitive
// os functions can be used with iterators to accomplish tasks in an efficient
// way. The use of iter.FlatMap with iterDir allows us to perform a depth-first
// search (limited to a depth of two in this case) without reading any more
// directory entries than necessary.
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"mtoohey.com/iter/v3"
)

func iterDir(dir string) iter.Iter[os.DirEntry] {
	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return func() (os.DirEntry, bool) {
		entries, err := f.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				return nil, false
			}

			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		return entries[0], true
	}
}

func isExecutableFile(entry os.DirEntry) bool {
	info, err := entry.Info()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return !info.IsDir() && info.Mode()&0o111 != 0
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s DIR\n", os.Args[0])
		os.Exit(1)
	}
	baseDir := os.Args[1]

	subdirs := iterDir(baseDir).Filter(os.DirEntry.IsDir)
	subdirPaths := iter.Map(subdirs, os.DirEntry.Name).Map(func(s string) string {
		return filepath.Join(baseDir, s)
	})
	firstExecutable, found := iter.FlatMap(subdirPaths, iterDir).Find(isExecutableFile)

	if found {
		fmt.Println(firstExecutable)
	}
}
