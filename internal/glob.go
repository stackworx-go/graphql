package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var path2regex = strings.NewReplacer(
	`.`, `\.`,
	`*`, `.+`,
	`\`, `[\\/]`,
	`/`, `[\\/]`,
)

// Glob Glob export
func Glob(f string) ([]string, error) {
	var matches []string
	var err error

	// for ** we want to override default globbing patterns and walk all
	// subdirectories to match schema files.
	if strings.Contains(f, "**") {
		pathParts := strings.SplitN(f, "**", 2)
		rest := strings.TrimPrefix(strings.TrimPrefix(pathParts[1], `\`), `/`)
		// turn the rest of the glob into a regex, anchored only at the end because ** allows
		// for any number of dirs in between and walk will let us match against the full path name
		globRe := regexp.MustCompile(path2regex.Replace(rest) + `$`)

		if err = filepath.Walk(pathParts[0], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if globRe.MatchString(strings.TrimPrefix(path, pathParts[0])) {
				matches = append(matches, path)
			}

			return nil
		}); err != nil {
			return nil, fmt.Errorf("failed to walk schema at root %s: %w", pathParts[0], err)
		}
	} else {
		matches, err = filepath.Glob(f)
		if err != nil {
			return nil, fmt.Errorf("failed to glob schema filename %s: %w", f, err)
		}
	}

	return matches, nil
}
