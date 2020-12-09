package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		fmt.Printf("%s\n", request.URL.Path)

		if _, ok := pathsToUrls[request.URL.Path]; !ok {
			fallback.ServeHTTP(response, request)
			return
		}

		http.Redirect(response, request, pathsToUrls[request.URL.Path], http.StatusMovedPermanently)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var dat []pathURL

	if err := yaml.Unmarshal(yml, &dat); err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(dat)

	var pathesMap = make(map[string]string)
	for _, path := range dat {
		pathesMap[path.Path] = path.URL
	}

	return MapHandler(pathesMap, fallback), nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
