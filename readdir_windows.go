package fs

import (
	"context"
	"path"
	"path/filepath"

	typ "github.com/aos-dev/go-storage/v2/types"
	"golang.org/x/sys/windows"
)

func (s *Storage) listDirNext(ctx context.Context, page *typ.ObjectPage) (err error) {
	input := page.Status.(*listDirInput)

	defer func() {
		// Make sure file has been close every time we return an error
		if err != nil && input.f != nil {
			_ = input.f.Close()
			input.f = nil
		}
	}()

	// Open dir before we read it.
	if input.f == nil {
		input.f, err = s.osOpen(input.rp)
		if err != nil {
			return
		}
	}

	// Every list dir will fetch 128 files.
	limit := 128
	var data windows.Win32finddata

	for limit > 0 {
		err = windows.FindNextFile(windows.Handle(input.f.Fd()), &data)
		// Whole dir has been read, return IterateDone to mark this iteration is done
		if err != nil && err == windows.ERROR_NO_MORE_FILES {
			return typ.IterateDone
		}
		if err != nil {
			return
		}

		name := windows.UTF16ToString(data.FileName[0:])
		if name == "." || name == ".." {
			continue
		}

		o := s.newObject(true)
		// Always keep service original name as ID.
		o.SetID(filepath.Join(input.rp, name))
		// Object's name should always be separated by slash (/)
		o.SetName(path.Join(input.dir, name))

		o.SetSize(int64(data.FileSizeHigh)<<32 + int64(data.FileSizeLow))

		switch {
		case data.FileAttributes&windows.FILE_ATTRIBUTE_DIRECTORY != 0:
			o.SetType(typ.ObjectTypeDir)
		case data.FileAttributes&windows.FILE_ATTRIBUTE_NORMAL != 0:
			o.SetType(typ.ObjectTypeFile)
		}
		page.Data = append(page.Data, o)

		limit--
	}
	return
}
