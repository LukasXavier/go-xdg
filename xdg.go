package xdg

import (
    "fmt"
    "os"
)

/*
XDG_DATA_HOME

Defines the base directory relative to which user-specific data files
should be stored.
*/
func DataHome() string {
    return format("XDG_DATA_HOME", ".local/share")
}

/*
XDG_CONFIG_HOME

Defines the base directory relative to which user-specific configuration
files should be stored.
*/
func ConfigHome() string {
    return format("XDG_CONFIG_HOME", ".config")
}

/*
XDG_STATE_HOME

contains state data that should persist between (application) restarts,
but that is not important or portable enough to the user that it should
be stored in $XDG_DATA_HOME.
It may contain:
  - actions history (logs, history recently used files, ...)
  - current state of the application that can be reused on a restart
    (view, layout, open files, undo history, ...)
*/
func StateHome() string {
    return format("XDG_STATE_HOME", ".local/state")
}

/*
XDG_CACHE_HOME

Defines the base directory relative to which user-specific non-essential
(cached) data should be written.
*/
func CacheHome() string {
    return format("XDG_CACHE_HOME", ".cache")
}

func format(env, path string) string {
    if path, found := os.LookupEnv(env); found {
    	return path
    }
    home, err := os.UserHomeDir()
    if err != nil {
    	panic("$HOME not set")
    }
    return fmt.Sprintf("%s/%s", home, path)
}
