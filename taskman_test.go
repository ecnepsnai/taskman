package taskman_test

import (
	"os"
	"testing"

	"github.com/ecnepsnai/logtic"
)

func TestMain(m *testing.M) {
	testSetup()
	code := m.Run()
	testShutdown()
	os.Exit(code)
}

func testSetup() {
	logtic.Log.FilePath = "/dev/null"
	logtic.Log.Level = logtic.LevelDebug
	logtic.Open()
}

func testShutdown() {
	logtic.Close()
}
