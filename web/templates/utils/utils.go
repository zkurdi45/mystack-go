package utils

import "fmt"

func GetStaticURL(path string, cacheBuster string) string {
	return fmt.Sprintf("/static%s?v=%s", path, cacheBuster)
}
