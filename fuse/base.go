package fuse

// DefaultFileSystem provides a filesystem that returns a suitable default for all methods.
// Most methods allow ENOSYS, which signals to FUSE that the operation is not implemented.
// Other methods simply return success, if the method is optional.
//
// This implementation is intended to be used as the base implementation for a filesystem, so that
// all methods not implemented by the derived type will be handled here.
//
// Usage eXAmple:
//   type MyFs struct {
//     fuse.DefaultFileSystem
//   }
type DefaultFileSystem struct{}

var _ FileSystem = &DefaultFileSystem{}

// Init implements FileSystem.
func (d *DefaultFileSystem) Init(*ConnInfo) {}

// Destroy implements FileSystem.
func (d *DefaultFileSystem) Destroy() {}

// StatFS implements FileSystem.
func (d *DefaultFileSystem) StatFS(ino InodeID) (*StatVFS, Status) {
	return nil, ENOSYS
}

// Lookup implements FileSystem.
func (d *DefaultFileSystem) Lookup(dir InodeID, name string) (entry *Entry, err Status) {
	return nil, ENOSYS
}

// Forget implements FileSystem.
func (d *DefaultFileSystem) Forget(ino InodeID, n int) {}

// Release implements FileSystem.
func (d *DefaultFileSystem) Release(ino InodeID, fi *FileInfo) Status {
	return ENOSYS
}

// ReleaseDir implements FileSystem.
func (d *DefaultFileSystem) ReleaseDir(ino InodeID, fi *FileInfo) Status {
	return ENOSYS
}

// FSync implements FileSystem.
func (d *DefaultFileSystem) FSync(ino InodeID, dataOnly bool, fi *FileInfo) Status {
	return ENOSYS
}

// FSyncDir implements FileSystem.
func (d *DefaultFileSystem) FSyncDir(ino InodeID, dataOnly bool, fi *FileInfo) Status {
	return ENOSYS
}

// Flush implements FileSystem.
func (d *DefaultFileSystem) Flush(ino InodeID, fi *FileInfo) Status {
	return ENOSYS
}

// GetAttr implements FileSystem.
func (d *DefaultFileSystem) GetAttr(ino InodeID, fi *FileInfo) (attr *InoAttr, err Status) {
	return nil, ENOSYS
}

// SetAttr implements FileSystem.
func (d *DefaultFileSystem) SetAttr(ino InodeID, attr *InoAttr, mask SetAttrMask, fi *FileInfo) (
	*InoAttr, Status) {
	return nil, ENOSYS
}

// ReadLink implements FileSystem.
func (d *DefaultFileSystem) ReadLink(ino InodeID) (string, Status) {
	return "", ENOSYS
}

// ReadDir implements FileSystem.
func (d *DefaultFileSystem) ReadDir(ino InodeID, fi *FileInfo, off int64, size int,
	w DirEntryWriter) Status {
	return ENOSYS
}

// Mknod implements FileSystem.
func (d *DefaultFileSystem) Mknod(p InodeID, name string, mode int, rdev int) (
	entry *Entry, err Status) {
	return nil, ENOSYS
}

// Access implements FileSystem.
func (d *DefaultFileSystem) Access(ino InodeID, mode int) Status {
	return ENOSYS
}

// Create implements FileSystem.
func (d *DefaultFileSystem) Create(p InodeID, name string, mode int, fi *FileInfo) (
	entry *Entry, err Status) {
	return nil, ENOSYS
}

// Open implements FileSystem.
func (d *DefaultFileSystem) Open(ino InodeID, fi *FileInfo) Status {
	return ENOSYS
}

// OpenDir implements FileSystem.
func (d *DefaultFileSystem) OpenDir(ino InodeID, fi *FileInfo) Status {
	return OK
}

// Read implements FileSystem.
func (d *DefaultFileSystem) Read(ino InodeID, size int64, off int64, fi *FileInfo) (
	data []byte, err Status) {
	return nil, ENOSYS
}

// Write implements FileSystem.
func (d *DefaultFileSystem) Write(p []byte, ino InodeID, off int64, fi *FileInfo) (
	n int, err Status) {
	return 0, ENOSYS
}

// Mkdir implements FileSystem.
func (d *DefaultFileSystem) Mkdir(p InodeID, name string, mode int) (
	entry *Entry, err Status) {
	return nil, ENOSYS
}

// Rmdir implements FileSystem.
func (d *DefaultFileSystem) Rmdir(p InodeID, name string) Status {
	return ENOSYS
}

// Symlink implements FileSystem.
func (d *DefaultFileSystem) Symlink(link string, p InodeID, name string) (*Entry, Status) {
	return nil, ENOSYS
}

// Link implements FileSystem.
func (d *DefaultFileSystem) Link(ino InodeID, newparent InodeID, name string) (*Entry, Status) {
	return nil, ENOSYS
}

// Rename implements FileSystem.
func (d *DefaultFileSystem) Rename(InodeID, string, InodeID, string) Status {
	return ENOSYS
}

// Unlink implements FileSystem.
func (d *DefaultFileSystem) Unlink(p InodeID, name string) Status {
	return ENOSYS
}

// ListXAttrs implements FileSystem.
func (d *DefaultFileSystem) ListXAttrs(ino InodeID) ([]string, Status) {
	return nil, ENOSYS
}

// GetXAttrSize implements FileSystem.
func (d *DefaultFileSystem) GetXAttrSize(ino InodeID, name string) (int, Status) {
	return 0, ENOSYS
}

// GetXAttr implements FileSystem.
func (d *DefaultFileSystem) GetXAttr(ino InodeID, name string, out []byte) (int, Status) {
	return 0, ENOSYS
}

// SetXAttr implements FileSystem.
func (d *DefaultFileSystem) SetXAttr(ino InodeID, name string, value []byte, flags int) Status {
	return ENOSYS
}

// RemoveXAttr implements FileSystem.
func (d *DefaultFileSystem) RemoveXAttr(ino InodeID, name string) Status {
	return ENOSYS
}
