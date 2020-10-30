package gwf

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_WithLogger(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	opt := WithLogger(logger)
	app := &Application{}
	opt.Apply(app)
	require.Equal(t, logger, app.logger)
}
