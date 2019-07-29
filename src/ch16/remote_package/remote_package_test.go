package remote_package

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(16)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
