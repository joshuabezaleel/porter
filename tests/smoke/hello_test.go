// +build smoke

package smoke

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloBundle(t *testing.T) {
	test, err := NewTest(t)
	defer test.Teardown()
	require.NoError(t, err, "test setup failed")

	test.RequirePorter("create")
	test.RequirePorter("build")

	ref := "localhost:5000/porter-hello:v0.1.1"
	test.RequirePorter("publish", "--reference", ref)

	test.RequirePorter("install", "--reference", ref)
	test.RequirePorter("installation", "show", "porter-hello")

	test.RequirePorter("upgrade")
	test.RequirePorter("installation", "show", "porter-hello")

	test.RequirePorter("uninstall")
	test.RequirePorter("installation", "show", "porter-hello")
}
