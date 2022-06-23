package loopback

import (
	"errors"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/vgough/go-fuse-c/fuse"
)

type InodeID = fuse.InodeID

type dir struct {
}

type file struct {
}

type FS struct {
	fuse.DefaultFileSystem

	root string

	mu     sync.RWMutex
	inodes map[InodeID]*inode
}

type inode struct {
	path string
	stat syscall.Stat_t
}

func (i *inode) Entry() *fuse.Entry {
	return &fuse.Entry{
		Ino:          InodeID(i.stat.Ino),
		Generation:   1, // TODO
		Attr:         i.Stat(),
		AttrTimeout:  1.0,
		EntryTimeout: 1.0,
	}
}

func (i *inode) Stat() *fuse.InoAttr {
	uid := int(i.stat.Uid)
	gid := int(i.stat.Gid)
	return &fuse.InoAttr{
		Ino:     InodeID(i.stat.Ino),
		Size:    i.stat.Size,
		Mode:    int(i.stat.Mode),
		NLink:   1,
		UID:     &uid,
		GID:     &gid,
		ATime:   time.Unix(int64(i.stat.Atim.Sec), int64(i.stat.Atim.Nsec)),
		CTime:   time.Unix(int64(i.stat.Ctim.Sec), int64(i.stat.Ctim.Nsec)),
		MTime:   time.Unix(int64(i.stat.Mtim.Sec), int64(i.stat.Mtim.Nsec)),
		Timeout: 1.0,
	}
}

var _ fuse.FileSystem = &FS{}

func New(root string) (*FS, error) {
	node := &inode{
		path: root,
	}
	if err := syscall.Stat(root, &node.stat); err != nil {
		return nil, err
	}
	return &FS{
		root: root,
		inodes: map[InodeID]*inode{
			1: node,
		},
	}, nil
}

// Init initializes a filesystem.
// Called before any other filesystem method.
func (fs *FS) Init(_ *fuse.ConnInfo) {
}

// Destroy cleans up a filesystem.
// Called on filesystem exit.
func (fs *FS) Destroy() {
}

func (fs *FS) getNode(id InodeID) *inode {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	return fs.inodes[id]
}

// // StatFS gets file system statistics.
// func (fs *FS) StatFS(ino InodeID) (*fuse.StatVFS, fuse.Status) {
// }

// Lookup finds a directory entry by name and get its attributes.
func (fs *FS) Lookup(dir InodeID, name string) (*fuse.Entry, fuse.Status) {
	parent := fs.getNode(dir)
	if parent == nil {
		return nil, fuse.ENOENT
	}
	if int(parent.stat.Mode)&syscall.S_IFDIR == 0 {
		return nil, fuse.ENOTDIR
	}

	fqn := filepath.Join(parent.path, name)
	node := &inode{
		path: fqn,
	}
	if err := syscall.Stat(fqn, &node.stat); err != nil {
		return nil, xlateErr(err)
	}

	fs.mu.Lock()
	fs.inodes[InodeID(node.stat.Ino)] = node
	fs.mu.Unlock()
	return node.Entry(), fuse.OK
}

func xlateErr(err error) fuse.Status {
	var errno syscall.Errno
	if errors.As(err, &errno) {
		return fuse.Status(errno)
	}

	return fuse.ENOSYS
}

// Forget limits the lifetime of an inode.
//
// The n parameter indicates the number of lookups previously performed on this inode.
// The filesystem may ignore forget calls if the inodes don't need to have a limited lifetime.
// On unmount it is not guaranteed that all reference dinodes will receive a forget message.
func (fs *FS) Forget(ino InodeID, n int) {
	panic("not implemented") // TODO: Implement
}

// Release drops an open file reference.
//
// Release is called when there are no more references to an open file: all file descriptors are
// closed and all memory mappings are unmapped.
//
// For every open call, there will be eXActly one release call.
//
// A filesystem may reply with an error, but error values are not returned to the close() or
// munmap() which triggered the release.
//
// fi.Handle will contain the value set by the open method, or will be undefined if the open
// method didn't set any value.
// fi.Flags will contain the same flags as for open.
func (fs *FS) Release(ino InodeID, fi *fuse.FileInfo) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// Flush is called on each close() of an opened file.
//
// Since file descriptors can be duplicated (dup, dup2, fork), for one open call there may be
// many flush calls.
//
// fi.Handle will contain the value set by the open method, or will be undefined if the open
// method didn't set any value.
//
// The name of the method is misleading. Unlike fsync, the filesystem is not forced to flush
// pending writes.
func (fs *FS) Flush(ino InodeID, fi *fuse.FileInfo) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// Fsync synchronizes file contents.
//
// If the dataOnly parameter is true, then only the user data should be flushed, not the
// metdata.
func (fs *FS) FSync(ino InodeID, dataOnly bool, fi *fuse.FileInfo) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// Getattr gets file attributes.
//
// fi is for future use, currently always nil.
func (fs *FS) GetAttr(ino InodeID, fi *fuse.FileInfo) (attr *fuse.InoAttr, err fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Setattr sets file attributes.
//
// In the 'attr' argument, only members indicated by the mask contain valid values.  Other
// members contain undefined values.
//
// If the setattr was invoked from the ftruncate() system call, the fi.Handle will contain the
// value set by the open method.  Otherwise, the fi argument may be nil.
func (fs *FS) SetAttr(ino InodeID, attr *fuse.InoAttr, mask fuse.SetAttrMask, fi *fuse.FileInfo) (*fuse.InoAttr, fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// ReadLink reads a symbolic link.
func (fs *FS) ReadLink(ino InodeID) (string, fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// ReadDir reads a directory.
//
// fi.Handle will contain the value set by the opendir method, or will be undefined if the
// opendir method didn't set any value.
//
// DirEntryWriter is used to add entries to the output buffer.
func (fs *FS) ReadDir(ino InodeID, fi *fuse.FileInfo, off int64, size int, w fuse.DirEntryWriter) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// OpenDir opens a directory.
//
// Filesystems may store an arbitrary file handle in fh.Handle and use this in other directory
// operations (ReadDir, ReleaseDir, FsyncDir). Filesystems may not store anything in fi.Handle,
// though that makes it impossible to implement standard conforming directory stream operations
// in case the contents of the directory can change between opendir and releasedir.
func (fs *FS) OpenDir(ino InodeID, fi *fuse.FileInfo) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// ReleaseDir drops an open file reference.
//
// For every OpenDir call, there will be eXActly one ReleaseDir call.
//
// fi.Handle will contain the value set by the OpenDir method, or will be undefined if the
// OpenDir method didn't set any value.
func (fs *FS) ReleaseDir(ino InodeID, fi *fuse.FileInfo) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// FsyncDir synchronizes directory contents.
//
// If the dataOnly parameter is true, then only the user data should be flushed, not the
// metdata.
func (fs *FS) FSyncDir(ino InodeID, dataOnly bool, fi *fuse.FileInfo) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// Mkdir creates a directory.
func (fs *FS) Mkdir(parent InodeID, name string, mode int) (*fuse.Entry, fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Rmdir removes a directory.
func (fs *FS) Rmdir(parent InodeID, name string) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// Rename renames a file or directory.
func (fs *FS) Rename(dir InodeID, name string, newdir InodeID, newname string) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// Symlink creates a symbolic link.
func (fs *FS) Symlink(link string, parent InodeID, name string) (*fuse.Entry, fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Link creates a hard link.
func (fs *FS) Link(ino InodeID, newparent InodeID, name string) (*fuse.Entry, fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Mknod creates a file node.
//
// This is used to create a regular file, or special files such as character devices, block
// devices, fifo or socket nodes.
func (fs *FS) Mknod(parent InodeID, name string, mode int, rdev int) (*fuse.Entry, fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Open makes a file available for read or write.
//
// Open flags are available in fi.Flags
//
// Filesystems may store an arbitrary file handle in fh.Handle and use this in other file
// operations (read, write, flush, release, fsync). Filesystems may also implement stateless file
// I/O and not store anything in fi.Handle.
func (fs *FS) Open(ino InodeID, fi *fuse.FileInfo) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// Read reads data from an open file.
//
// Read should return eXActly the number of bytes requested except on EOF or error.
//
// fi.Handle will contain the value set by the open method, if any.
func (fs *FS) Read(ino InodeID, size int64, off int64, fi *fuse.FileInfo) (data []byte, err fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Write writes data to an open file.
//
// Write should return eXActly the number of bytes requested except on error.
//
// fi.handle will contain the value set by the open method, if any.
func (fs *FS) Write(p []byte, ino InodeID, off int64, fi *fuse.FileInfo) (n int, err fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Unlink removes a file.
func (fs *FS) Unlink(parent InodeID, name string) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// Access checks file access permissions.
//
// This will be called for the access() system call.  If the 'default_permissions' mount option
// is given, this method is not called.
func (fs *FS) Access(ino InodeID, mask int) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// Create creates and opens a file.
//
// If the file does not exist, first create it with the specified mode and then open it.
//
// Open flags are available in fi.Flags.
//
// Filesystems may store an arbitrary file handle in fi.Handle and use this in all other file
// operations (Read, Write, Flush, Release, FSync).
//
// If this method is not implemented, then Mknod and Open methods will be called instead.
func (fs *FS) Create(parent InodeID, name string, mode int, fi *fuse.FileInfo) (*fuse.Entry, fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Returns a list of the extended attribute keys.
func (fs *FS) ListXAttrs(ino InodeID) ([]string, fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Returns the size of the attribute value.
func (fs *FS) GetXAttrSize(ino InodeID, name string) (int, fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Get an extended attribute.
// Result placed in out buffer.
// Returns the number of bytes copied.
func (fs *FS) GetXAttr(ino InodeID, name string, out []byte) (int, fuse.Status) {
	panic("not implemented") // TODO: Implement
}

// Set an extended attribute.
func (fs *FS) SetXAttr(ino InodeID, name string, value []byte, flags int) fuse.Status {
	panic("not implemented") // TODO: Implement
}

// Remove an extended attribute.
func (fs *FS) RemoveXAttr(ino InodeID, name string) fuse.Status {
	panic("not implemented") // TODO: Implement
}
