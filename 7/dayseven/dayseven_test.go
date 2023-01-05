package dayseven

import (
	"testing"
)

const (
	testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`
)

func TestNewFiles(t *testing.T) {
	t.Parallel()
	var f File
	fileTests := []struct {
		name     string
		filename string
		size     int
		expected File
	}{
		{"regular file", "plain", 123, File{"plain", 123}},
		{"file with extension", "ext.ext", 234, File{"ext.ext", 234}},
		{"file with spaces", "spaces in the name", 345, File{"spaces in the name", 345}},
	}
	for _, ft := range fileTests {
		t.Run(ft.name, func(t *testing.T) {
			t.Parallel()
			if f = NewFile(ft.filename, ft.size); f != ft.expected || f.IsDir() || !f.IsFile() {
				t.Errorf("New file failure. Got %s, wanted file name of %q with size %d.", f, ft.expected.Name(), ft.expected.Size())
			}
		})
	}
}

func TestNewDirectories(t *testing.T) {
	t.Parallel()
	emptyDir := NewDirectory("emptyDir")
	dirWithFiles := NewDirectory("withFiles", NewFile("onefile", 1000), NewFile("otherfile", 100))
	dirWithNestedDir := NewDirectory("withNested", NewDirectory("nested", NewFile("afile", 100)))
	dirTests := []struct {
		name         string
		test         Directory
		expectedName string
		expectedSize int
	}{
		{"empty directory", emptyDir, "emptyDir", 0},
		{"directory with two files", dirWithFiles, "withFiles", 1100},
		{"nested directory", dirWithNestedDir, "withNested", 100},
	}
	for _, dt := range dirTests {
		t.Run(dt.name, func(t *testing.T) {
			t.Parallel()
			if dt.test.Name() != dt.expectedName || dt.test.Size() != dt.expectedSize || dt.test.IsFile() || !dt.test.IsDir() {
				t.Errorf("ERROR: Got %s, wanted dir name of %q with size %d.", dt.test, dt.expectedName, dt.expectedSize)
			}
		})
	}
}

func TestDirectoryAssociation(t *testing.T) {
	t.Parallel()

	// Directories
	rootDir := NewDirectory("/")
	emptyDir := NewDirectory("emptyDir")
	dirWithOneFile := NewDirectory("withBeegFile", NewFile("beegfile", 1000))
	deepNest := NewDirectory("deeplyNested", NewFile("nestedFile", 5))

	// Adding files with AddTo()
	dirWithNested := NewDirectory("withFiles")
	NewFile("onefile", 100).AddTo(&dirWithNested)
	NewFile("otherfile", 10).AddTo(&dirWithNested)

	// Associating directories:
	// - root
	//   - emptyDir
	//   - dirWithOneFile
	//   - dirWithTwoFiles
	//     - deepNest
	emptyDir.AddParent(&rootDir)
	dirWithOneFile.AddParent(&rootDir)
	dirWithNested.AddParent(&rootDir)
	deepNest.AddParent(&dirWithOneFile)

	expectedSize := emptyDir.Size() + dirWithOneFile.Size() + dirWithNested.Size()

	if rootDir.Size() != expectedSize {
		t.Errorf("ERROR: Incorrect size for root directory. Got %d, wanted %d", rootDir.Size(), expectedSize)
	}
}