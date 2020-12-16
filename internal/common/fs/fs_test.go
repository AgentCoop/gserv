package fs_test

import (
	"github.com/AgentCoop/gserv/internal/common/fs"
	"testing"
)

func TestScanner_Run(t *testing.T) {
	scanner := fs.NewScanner(".")
	_ := scanner.Run()
}
