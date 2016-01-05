package test

import (
	"fmt"
	"os"
	"steamtest/src/config"
	"steamtest/src/constants"
)

// SetupEnvironment sets up the environment for tests. This should only be
// called once per package and only in the first _test.go file of the package
// that needs it.
func SetupEnvironment() {
	fmt.Println("Setting up test environment...")
	// Need base directory for config and other files
	err := os.Chdir("../..")
	if err != nil {
		panic("Unable to change directory for tests")
	}
	// Remove old test files
	deleteFiles(constants.TestTempDirectory,
		constants.DumpFileFullPath(
			config.ReadConfig().DebugConfig.ServerDumpFilename))

	// use testing configuration
	config.CreateTestConfig()
	constants.IsTest = true
}

func deleteFiles(filepaths ...string) {
	fmt.Println("Running pre-test cleanup...")
	for _, fps := range filepaths {
		err := os.RemoveAll(fps)
		if err != nil {
			fmt.Printf("Error running test cleanup; unable to remove %s: %s", fps, err)
		}
	}
}
