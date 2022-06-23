package fuse

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var fs FileSystem
var fsID int

func TestMain(m *testing.M) {
	EnableBridgeTestMode()

	fs = NewMemFS()
	fsID = RegisterFS(fs)

	r := m.Run()

	DeregisterFS(fsID)
	os.Exit(r)
}

func TestVersion(t *testing.T) {
	version := Version()
	require.True(t, version >= 26)
}

func TestLookup(t *testing.T) {
	fileEnt, _ := fs.Mknod(1, "exists", 0444, 0)

	t.Run("Lookup invalid inode", func(t *testing.T) {
		bridgeLookup(fsID, 1000, "test", func(id int, r interface{}) int {
			switch r := r.(type) {
			case *replyErr:
				require.Equal(t, ENOENT, r.err)

			default:
				t.Errorf("Unexpected reply: %#v", r)
			}
			return int(OK)
		})
	})

	t.Run("Lookup invalid file", func(t *testing.T) {
		bridgeLookup(fsID, 1, "test", func(id int, r interface{}) int {
			switch r := r.(type) {
			case *replyErr:
				require.Equal(t, ENOENT, r.err)

			default:
				t.Errorf("Unexpected reply: %#v", r)
			}

			return int(OK)
		})
	})

	t.Run("Lookup valid file", func(t *testing.T) {
		bridgeLookup(fsID, 1, "exists", func(id int, r interface{}) int {
			require.IsType(t, &replyEntry{}, r)
			return int(OK)
		})
	})

	t.Run("Lookup invalid node type", func(t *testing.T) {
		// Pass a file inode as the directory.
		bridgeLookup(fsID, fileEnt.Ino, "test", func(id int, r interface{}) int {
			switch r := r.(type) {
			case *replyErr:
				require.Equal(t, ENOTDIR, r.err)

			default:
				t.Errorf("Unexpected reply: %#v", r)
			}
			return int(OK)
		})
	})
}

func TestForget(t *testing.T) {
	t.Run("Forget uses reply_none", func(t *testing.T) {
		bridgeForget(fsID, 100, 1, func(id int, r interface{}) int {
			require.IsType(t, &replyNone{}, r)
			return int(OK)
		})
	})
}

func TestGetAttr(t *testing.T) {
	t.Run("GetAttr on existing directory", func(t *testing.T) {
		bridgeGetAttr(fsID, 1, func(id int, r interface{}) int {
			require.IsType(t, &replyAttr{}, r)
			a, _ := r.(*replyAttr)
			stat := a.attr
			require.NotNil(t, stat)
			require.EqualValues(t, 1, stat.st_ino)
			return int(OK)
		})
	})

	t.Run("GetAttr on nonexistant node", func(t *testing.T) {
		bridgeGetAttr(fsID, 1000, func(id int, r interface{}) int {
			switch r := r.(type) {
			case *replyErr:
				require.Equal(t, ENOENT, r.err)

			default:
				t.Errorf("Unexpected reply: %#v", r)
			}
			return int(OK)
		})
	})
}

func TestStatFs(t *testing.T) {
	t.Run("StatFs on undefined inode", func(t *testing.T) {
		bridgeStatFs(fsID, 0, func(id int, r interface{}) int {
			require.IsType(t, &replyStatFs{}, r)
			return int(OK)
		})
	})
}
