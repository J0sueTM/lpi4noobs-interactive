package artifact

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Document struct {
	Name     string     `json:"name"`
	Folder   string     `json:"folder"`
	Filename string     `json:"filename"`
	Children []Document `json:"children"`
}

type Content struct {
	BaseURL   string     `json:"base_url"`
	Documents []Document `json:"documents"`
}

const (
	resourcesDir string = "resources"
	cacheDir     string = "cache"
)

func NewContent(filename string) (Artifact, error) {
	filePath := fmt.Sprintf("%s/%s", resourcesDir, filename)
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	ctt := Content{}
	err = json.Unmarshal([]byte(file), &ctt)
	if err != nil {
		return nil, err
	}

	if !ctt.HasCache(cacheDir) {
		err := os.Mkdir(cacheDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	// TODO: Move back to checking after implemented
	ctt.FetchCache()

	ctt.LoadCache()

	return &ctt, nil
}

func (*Content) HasCache(dirPath string) bool {
	_, err := os.Stat(dirPath)
	return !errors.Is(err, os.ErrNotExist)
}

func (*Content) LoadCache() error {
	// TODO: Implement me

	return nil
}

func (*Content) WriteCache() error {
	// TODO: Implement me
	return nil
}

func (ctt *Content) FetchCache() error {
	// TODO: Implement me
	for _, doc := range ctt.Documents {
		fetchDocument(&doc)
	}

	return nil
}

func fetchDocument(doc *Document) error {
	// TODO: Implement me

	for _, cDoc := range doc.Children {
		return fetchDocument(&cDoc)
	}

	return nil
}
