package version

import "testing"

func TestValidSimpleVersionFormatSnapshot(t *testing.T) {
	for _, value := range []string{"1.0-snapshot", "1.0-SNAPSHOT", "1.0-SNAPSHOT1"} {
		if !ValidSimpleVersionFormat(value) {
			t.Errorf("ValidSimpleVersionFormat(%q) rejected a supported snapshot version", value)
		}
		if !Compare(value, "1.0-dev", "<") {
			t.Errorf("Compare(%q, 1.0-dev) did not preserve snapshot ordering", value)
		}
	}
	for _, value := range []string{"1.0-unknown", "1.0-snapshot.unknown", "1.0-SNAPSHOT.unknown"} {
		if ValidSimpleVersionFormat(value) {
			t.Errorf("ValidSimpleVersionFormat(%q) accepted an unknown component", value)
		}
	}
}
