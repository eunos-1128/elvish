package runtime

import (
	"embed"
	"errors"
	"testing"

	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/evaltest"
	"src.elv.sh/pkg/testutil"
)

//go:embed *.elvts
var transcripts embed.FS

func TestTranscripts(t *testing.T) {
	evaltest.TestTranscriptsInFS(t, transcripts,
		"use-runtime-good-paths", func(t *testing.T, ev *eval.Evaler) {
			testutil.Set(t, &osExecutable,
				func() (string, error) { return "/path/to/elvish", nil })
			ev.LibDirs = []string{"/lib/1", "/lib/2"}
			ev.RcPath = "/path/to/rc.elv"
			ev.EffectiveRcPath = "/path/to/effective/rc.elv"

			ev.ExtendGlobal(eval.BuildNs().AddNs("runtime", Ns(ev)))
		},
		"use-runtime-bad-paths", func(t *testing.T, ev *eval.Evaler) {
			testutil.Set(t, &osExecutable,
				func() (string, error) { return "bad", errors.New("bad") })
			// Leaving all the path fields in ev unset

			ev.ExtendGlobal(eval.BuildNs().AddNs("runtime", Ns(ev)))
		},
	)
}
