package main

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathsToUrls []struct {
		Path string `yaml:"path"`
		URL  string `yaml:"url"`
	}

	err := yaml.Unmarshal(yml, &pathsToUrls)
	if err != nil {
		return nil, err
	}

	pathsMap := make(map[string]string)
	for _, item := range pathsToUrls {
		pathsMap[item.Path] = item.URL
	}

	return MapHandler(pathsMap, fallback), nil
}
