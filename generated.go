// Code generated by go generate via internal/cmd/service; DO NOT EDIT.
package fs

import (
	"context"
	"io"

	ps "github.com/aos-dev/go-storage/v2/pairs"
	"github.com/aos-dev/go-storage/v2/pkg/credential"
	"github.com/aos-dev/go-storage/v2/pkg/endpoint"
	"github.com/aos-dev/go-storage/v2/pkg/httpclient"
	"github.com/aos-dev/go-storage/v2/services"
	. "github.com/aos-dev/go-storage/v2/types"
)

var _ credential.Provider
var _ endpoint.Provider
var _ Storager
var _ services.ServiceError
var _ httpclient.Options

// Type is the type for fs
const Type = "fs"

// Service available pairs.
const (
	// EnableLinkFollow will // EnableLinkFollow
	PairEnableLinkFollow = "fs_enable_link_follow"
)

// Service available infos.
const ()

// WithEnableLinkFollow will apply enable_link_follow value to Options
// This pair is used to // EnableLinkFollow
func WithEnableLinkFollow(v bool) Pair {
	return Pair{
		Key:   PairEnableLinkFollow,
		Value: v,
	}
}

// pairStorageNewMap holds all available pairs
var pairStorageNewMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	ps.WorkDir: struct{}{},
	// Generated pairs
	ps.HTTPClientOptions: struct{}{},
}

// pairStorageNew is the parsed struct
type pairStorageNew struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	HasWorkDir bool
	WorkDir    string
	// Generated pairs
	HasHTTPClientOptions bool
	HTTPClientOptions    *httpclient.Options
}

// parsePairStorageNew will parse *Pair slice into *pairStorageNew
func parsePairStorageNew(opts []Pair) (*pairStorageNew, error) {
	result := &pairStorageNew{
		pairs: opts,
	}

	values := make(map[string]interface{})
	for _, v := range opts {
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool

	// Handle required pairs
	// Handle optional pairs
	v, ok = values[ps.WorkDir]
	if ok {
		result.HasWorkDir = true
		result.WorkDir = v.(string)
	}
	// Handle generated pairs
	v, ok = values[ps.HTTPClientOptions]
	if ok {
		result.HasHTTPClientOptions = true
		result.HTTPClientOptions = v.(*httpclient.Options)
	}

	return result, nil
}

// pairStorageCopyMap holds all available pairs
var pairStorageCopyMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	// Generated pairs
}

// pairStorageCopy is the parsed struct
type pairStorageCopy struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairStorageCopy will parse *Pair slice into *pairStorageCopy
func parsePairStorageCopy(opts []Pair) (*pairStorageCopy, error) {
	result := &pairStorageCopy{
		pairs: opts,
	}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageCopyMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	// Handle required pairs
	// Handle optional pairs
	// Handle generated pairs

	return result, nil
}

// pairStorageDeleteMap holds all available pairs
var pairStorageDeleteMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	// Generated pairs
}

// pairStorageDelete is the parsed struct
type pairStorageDelete struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairStorageDelete will parse *Pair slice into *pairStorageDelete
func parsePairStorageDelete(opts []Pair) (*pairStorageDelete, error) {
	result := &pairStorageDelete{
		pairs: opts,
	}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageDeleteMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	// Handle required pairs
	// Handle optional pairs
	// Handle generated pairs

	return result, nil
}

// pairStorageListDirMap holds all available pairs
var pairStorageListDirMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	ps.ContinuationToken: struct{}{},
	PairEnableLinkFollow: struct{}{},
	// Generated pairs
}

// pairStorageListDir is the parsed struct
type pairStorageListDir struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	HasContinuationToken bool
	ContinuationToken    string
	HasEnableLinkFollow  bool
	EnableLinkFollow     bool
	// Generated pairs
}

// parsePairStorageListDir will parse *Pair slice into *pairStorageListDir
func parsePairStorageListDir(opts []Pair) (*pairStorageListDir, error) {
	result := &pairStorageListDir{
		pairs: opts,
	}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageListDirMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool

	// Handle required pairs
	// Handle optional pairs
	v, ok = values[ps.ContinuationToken]
	if ok {
		result.HasContinuationToken = true
		result.ContinuationToken = v.(string)
	}
	v, ok = values[PairEnableLinkFollow]
	if ok {
		result.HasEnableLinkFollow = true
		result.EnableLinkFollow = v.(bool)
	}
	// Handle generated pairs

	return result, nil
}

// pairStorageMetadataMap holds all available pairs
var pairStorageMetadataMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	// Generated pairs
}

// pairStorageMetadata is the parsed struct
type pairStorageMetadata struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairStorageMetadata will parse *Pair slice into *pairStorageMetadata
func parsePairStorageMetadata(opts []Pair) (*pairStorageMetadata, error) {
	result := &pairStorageMetadata{
		pairs: opts,
	}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageMetadataMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	// Handle required pairs
	// Handle optional pairs
	// Handle generated pairs

	return result, nil
}

// pairStorageMoveMap holds all available pairs
var pairStorageMoveMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	// Generated pairs
}

// pairStorageMove is the parsed struct
type pairStorageMove struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairStorageMove will parse *Pair slice into *pairStorageMove
func parsePairStorageMove(opts []Pair) (*pairStorageMove, error) {
	result := &pairStorageMove{
		pairs: opts,
	}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageMoveMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	// Handle required pairs
	// Handle optional pairs
	// Handle generated pairs

	return result, nil
}

// pairStorageReadMap holds all available pairs
var pairStorageReadMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	ps.Offset:           struct{}{},
	ps.ReadCallbackFunc: struct{}{},
	ps.Size:             struct{}{},
	// Generated pairs
}

// pairStorageRead is the parsed struct
type pairStorageRead struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	HasOffset           bool
	Offset              int64
	HasReadCallbackFunc bool
	ReadCallbackFunc    func([]byte)
	HasSize             bool
	Size                int64
	// Generated pairs
}

// parsePairStorageRead will parse *Pair slice into *pairStorageRead
func parsePairStorageRead(opts []Pair) (*pairStorageRead, error) {
	result := &pairStorageRead{
		pairs: opts,
	}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageReadMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool

	// Handle required pairs
	// Handle optional pairs
	v, ok = values[ps.Offset]
	if ok {
		result.HasOffset = true
		result.Offset = v.(int64)
	}
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}
	v, ok = values[ps.Size]
	if ok {
		result.HasSize = true
		result.Size = v.(int64)
	}
	// Handle generated pairs

	return result, nil
}

// pairStorageStatMap holds all available pairs
var pairStorageStatMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	// Generated pairs
}

// pairStorageStat is the parsed struct
type pairStorageStat struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	// Generated pairs
}

// parsePairStorageStat will parse *Pair slice into *pairStorageStat
func parsePairStorageStat(opts []Pair) (*pairStorageStat, error) {
	result := &pairStorageStat{
		pairs: opts,
	}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageStatMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}

	// Handle required pairs
	// Handle optional pairs
	// Handle generated pairs

	return result, nil
}

// pairStorageWriteMap holds all available pairs
var pairStorageWriteMap = map[string]struct{}{
	// Required pairs
	// Optional pairs
	ps.Size: struct{}{},
	// Generated pairs
	ps.ReadCallbackFunc: struct{}{},
}

// pairStorageWrite is the parsed struct
type pairStorageWrite struct {
	pairs []Pair

	// Required pairs
	// Optional pairs
	HasSize bool
	Size    int64
	// Generated pairs
	HasReadCallbackFunc bool
	ReadCallbackFunc    func([]byte)
}

// parsePairStorageWrite will parse *Pair slice into *pairStorageWrite
func parsePairStorageWrite(opts []Pair) (*pairStorageWrite, error) {
	result := &pairStorageWrite{
		pairs: opts,
	}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := pairStorageWriteMap[v.Key]; !ok {
			return nil, services.NewPairUnsupportedError(v)
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool

	// Handle required pairs
	// Handle optional pairs
	v, ok = values[ps.Size]
	if ok {
		result.HasSize = true
		result.Size = v.(int64)
	}
	// Handle generated pairs
	v, ok = values[ps.ReadCallbackFunc]
	if ok {
		result.HasReadCallbackFunc = true
		result.ReadCallbackFunc = v.(func([]byte))
	}

	return result, nil
}

// Copy will copy an Object or multiple object in the service.
//
// This function will create a context by default.
func (s *Storage) Copy(src string, dst string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.CopyWithContext(ctx, src, dst, pairs...)
}

// CopyWithContext will copy an Object or multiple object in the service.
func (s *Storage) CopyWithContext(ctx context.Context, src string, dst string, pairs ...Pair) (err error) {
	defer func() {
		err = s.formatError("copy", err, src, dst)
	}()
	var opt *pairStorageCopy
	opt, err = parsePairStorageCopy(pairs)
	if err != nil {
		return
	}

	return s.copy(ctx, src, dst, opt)
}

// Delete will delete an Object from service.
//
// This function will create a context by default.
func (s *Storage) Delete(path string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.DeleteWithContext(ctx, path, pairs...)
}

// DeleteWithContext will delete an Object from service.
func (s *Storage) DeleteWithContext(ctx context.Context, path string, pairs ...Pair) (err error) {
	defer func() {
		err = s.formatError("delete", err, path)
	}()
	var opt *pairStorageDelete
	opt, err = parsePairStorageDelete(pairs)
	if err != nil {
		return
	}

	return s.delete(ctx, path, opt)
}

// ListDir will return list a specific dir.
//
// This function will create a context by default.
func (s *Storage) ListDir(dir string, pairs ...Pair) (oi *ObjectIterator, err error) {
	ctx := context.Background()
	return s.ListDirWithContext(ctx, dir, pairs...)
}

// ListDirWithContext will return list a specific dir.
func (s *Storage) ListDirWithContext(ctx context.Context, dir string, pairs ...Pair) (oi *ObjectIterator, err error) {
	defer func() {
		err = s.formatError("list_dir", err, dir)
	}()
	var opt *pairStorageListDir
	opt, err = parsePairStorageListDir(pairs)
	if err != nil {
		return
	}

	return s.listDir(ctx, dir, opt)
}

// Metadata will return current storager's metadata.
//
// This function will create a context by default.
func (s *Storage) Metadata(pairs ...Pair) (meta StorageMeta, err error) {
	ctx := context.Background()
	return s.MetadataWithContext(ctx, pairs...)
}

// MetadataWithContext will return current storager's metadata.
func (s *Storage) MetadataWithContext(ctx context.Context, pairs ...Pair) (meta StorageMeta, err error) {
	defer func() {
		err = s.formatError("metadata", err)
	}()
	var opt *pairStorageMetadata
	opt, err = parsePairStorageMetadata(pairs)
	if err != nil {
		return
	}

	return s.metadata(ctx, opt)
}

// Move will move an object in the service.
//
// This function will create a context by default.
func (s *Storage) Move(src string, dst string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.MoveWithContext(ctx, src, dst, pairs...)
}

// MoveWithContext will move an object in the service.
func (s *Storage) MoveWithContext(ctx context.Context, src string, dst string, pairs ...Pair) (err error) {
	defer func() {
		err = s.formatError("move", err, src, dst)
	}()
	var opt *pairStorageMove
	opt, err = parsePairStorageMove(pairs)
	if err != nil {
		return
	}

	return s.move(ctx, src, dst, opt)
}

// Read will read the file's data.
//
// This function will create a context by default.
func (s *Storage) Read(path string, w io.Writer, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.ReadWithContext(ctx, path, w, pairs...)
}

// ReadWithContext will read the file's data.
func (s *Storage) ReadWithContext(ctx context.Context, path string, w io.Writer, pairs ...Pair) (n int64, err error) {
	defer func() {
		err = s.formatError("read", err, path)
	}()
	var opt *pairStorageRead
	opt, err = parsePairStorageRead(pairs)
	if err != nil {
		return
	}

	return s.read(ctx, path, w, opt)
}

// Stat will stat a path to get info of an object.
//
// This function will create a context by default.
func (s *Storage) Stat(path string, pairs ...Pair) (o *Object, err error) {
	ctx := context.Background()
	return s.StatWithContext(ctx, path, pairs...)
}

// StatWithContext will stat a path to get info of an object.
func (s *Storage) StatWithContext(ctx context.Context, path string, pairs ...Pair) (o *Object, err error) {
	defer func() {
		err = s.formatError("stat", err, path)
	}()
	var opt *pairStorageStat
	opt, err = parsePairStorageStat(pairs)
	if err != nil {
		return
	}

	return s.stat(ctx, path, opt)
}

// Write will write data into a file.
//
// This function will create a context by default.
func (s *Storage) Write(path string, r io.Reader, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.WriteWithContext(ctx, path, r, pairs...)
}

// WriteWithContext will write data into a file.
func (s *Storage) WriteWithContext(ctx context.Context, path string, r io.Reader, pairs ...Pair) (n int64, err error) {
	defer func() {
		err = s.formatError("write", err, path)
	}()
	var opt *pairStorageWrite
	opt, err = parsePairStorageWrite(pairs)
	if err != nil {
		return
	}

	return s.write(ctx, path, r, opt)
}
