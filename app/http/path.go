package http

import (
	"errors"
	"fmt"
	"strings"
)

var PathNotFoundError = errors.New("path not found")
var EmptyPathError = errors.New("path cannot be empty")

func validatePath(path string) error {
	if path == "" {
		return fmt.Errorf("%w: %s", EmptyPathError, path)
	}
	if path != "/" {
		return fmt.Errorf("%w: %s", PathNotFoundError, path)
	}

	return nil
}

type Path string

const EchoPath Path = "/echo"
const UserAgentPath Path = "/user-agent"
const RootPath Path = "/"

var paths map[string]Path = map[string]Path{"echo": EchoPath, "user-agent": UserAgentPath}

func identifyPath(path string) (Path, string, error) {
	if path == string(RootPath) {
		fmt.Printf("found root path")
		return RootPath, "", nil
	}

	if len(path) < 1 {
		fmt.Printf("found empty path")
		return "", "", fmt.Errorf("%w: %s", EmptyPathError, path)
	}

	pathParts := strings.Split(path, "/")
	if len(pathParts) < 1 {
		fmt.Printf("found empty path (with slash)")
		return "", "", fmt.Errorf("%w: %s", EmptyPathError, path)
	}
	pathParts = pathParts[1:]

	realPath, ok := paths[pathParts[0]]
	if !ok {
		fmt.Printf("path not found: %s; rest: %+v, %d", pathParts[0], pathParts, len(pathParts))
		return "", "", fmt.Errorf("%w: %s", PathNotFoundError, path)
	}

	return realPath, strings.Join(pathParts[1:], "/"), nil
}
