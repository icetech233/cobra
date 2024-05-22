// Copyright 2024

package cobra

import (
	"fmt"
	"os"
)

const (
	activeHelpMarker = "_activeHelp_ "
	// The below values should not be changed: programs will be using them explicitly
	// in their user documentation, and users will be using them explicitly.
	activeHelpEnvVarSuffix  = "ACTIVE_HELP"
	activeHelpGlobalEnvVar  = configEnvVarGlobalPrefix + "_" + activeHelpEnvVarSuffix
	activeHelpGlobalDisable = "0"
)

// AppendActiveHelp adds the specified string to the specified array to be used as ActiveHelp.
// Such strings will be processed by the completion script and will be shown as ActiveHelp
// to the user.
// The array parameter should be the array that will contain the completions.
// This function can be called multiple times before and/or after completions are added to
// the array.  Each time this function is called with the same array, the new
// ActiveHelp line will be shown below the previous ones when completion is triggered.
func AppendActiveHelp(compArray []string, activeHelpStr string) []string {
	return append(compArray, fmt.Sprintf("%s%s", activeHelpMarker, activeHelpStr))
}

// GetActiveHelpConfig returns the value of the ActiveHelp environment variable
// <PROGRAM>_ACTIVE_HELP where <PROGRAM> is the name of the root command in upper
// case, with all non-ASCII-alphanumeric characters replaced by `_`.
// It will always return "0" if the global environment variable COBRA_ACTIVE_HELP
// is set to "0".
func GetActiveHelpConfig(cmd *Command) string {
	activeHelpCfg := os.Getenv(activeHelpGlobalEnvVar)
	if activeHelpCfg != activeHelpGlobalDisable {
		activeHelpCfg = os.Getenv(activeHelpEnvVar(cmd.Root().Name()))
	}
	return activeHelpCfg
}

// activeHelpEnvVar returns the name of the program-specific ActiveHelp environment
// variable.  It has the format <PROGRAM>_ACTIVE_HELP where <PROGRAM> is the name of the
// root command in upper case, with all non-ASCII-alphanumeric characters replaced by `_`.
func activeHelpEnvVar(name string) string {
	return configEnvVar(name, activeHelpEnvVarSuffix)
}