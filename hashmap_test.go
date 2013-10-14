package hashmap

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	kv := map[string]interface{}{
		"foo": "bar",
		"baz": 1,
	}

	hm := NewHashMap()
	hm.SetFromMap(kv)

	// String()
	hmstr := hm.String()

	if hmstr != `baz=1 foo="bar"` {
		t.Errorf("HashMap.String(): expected %q  got: %q", `baz=1 foo="bar"`,
			hmstr)
	}

	// Get()
	if hm.Get("foo") != "bar" {
		t.Errorf("HashMap[foo]: expected: %q  got: %q", "bar", hm.Get("foo"))
	}

	if hm.Get("baz") != 1 {
		t.Errorf("HashMap[baz]: expected: %d  got: %q", 1, hm.Get("baz"))
	}

	// Set()
	hm.Set("qux", "mock")

	if hm.kv["qux"] != "mock" {
		t.Errorf(`HashMap.Set("qux") expected: %q  got: %q`, "mock",
			hm.kv["qux"])
	}

	// Test that String representation of HashMap replaces spaces in keys
	h2 := NewHashMap()
	h2.Set("test key", "test value")
	hmstr = h2.String()

	if hmstr != `test_key="test value"` {
		t.Errorf("HashMap.encodeKeys: expected: %q  got: %q",
			`test_key="test_value"`, hmstr)
	}
}
