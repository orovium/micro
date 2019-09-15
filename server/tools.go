package server

import "os"

func envExist(envVar string) bool {
	return os.Getenv(envVar) != ""
}
