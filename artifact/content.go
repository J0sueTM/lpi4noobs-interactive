package artifact

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type Document struct {
	Name     string     `json:"name"`
	Folder   string     `json:"folder"`
	Filename string     `json:"filename"`
	Children []Document `json:"children"`
	Body     []byte
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
	filePath := path.Clean(fmt.Sprintf("%s/%s", resourcesDir, filename))
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	ctt := Content{}
	err = json.Unmarshal([]byte(file), &ctt)
	if err != nil {
		return nil, err
	}

	cttCacheDir := path.Clean(fmt.Sprintf("%s/content", cacheDir))
	if !ctt.HasCache(cttCacheDir) {
		err := os.MkdirAll(cttCacheDir, os.ModePerm)
		if err != nil {
			return nil, err
		}

		err = ctt.FetchCache()
		if err != nil {
			return nil, err
		}

		err = ctt.WriteCache()
		if err != nil {
			return nil, err
		}
	}

	err = ctt.LoadCache()
	if err != nil {
		return nil, err
	}

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

func (ctt *Content) WriteCache() error {
	fmt.Println("writing cache...")

	baseCacheDir := path.Clean(fmt.Sprintf("%s/content", cacheDir))
	return writeContent(baseCacheDir, &ctt.Documents[0])
}

func writeContent(baseDir string, parentDoc *Document) error {
	docDir := path.Clean(fmt.Sprintf("%s/%s", baseDir, parentDoc.Folder))
	_, err := os.Stat(docDir)
	if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(docDir, os.ModePerm)
	}

	docFilePath := path.Clean(fmt.Sprintf("%s/%s", docDir, parentDoc.Filename))
	fmt.Printf("writing %s\n", docFilePath)

	err = os.WriteFile(docFilePath, parentDoc.Body, os.ModePerm)
	if err != nil {
		return err
	}

	// using `i` instead of `_, doc` for mutability
	for i := range parentDoc.Children {
		err = writeContent(docDir, &parentDoc.Children[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (ctt *Content) FetchCache() error {
	fmt.Println("fetching cache...")

	return fetchContent(ctt.BaseURL, &ctt.Documents[0])
}

func fetchContent(baseURL string, parentDoc *Document) error {
	docBody, err := fetchDocument(baseURL, parentDoc)
	if err != nil {
		return err
	}
	parentDoc.Body = docBody

	// again, using `i` instead of `_, doc` for mutability
	for i := range parentDoc.Children {
		curDoc := &parentDoc.Children[i]
		curBaseURL := fmt.Sprintf("%s%s", baseURL, parentDoc.Folder)

		err = fetchContent(curBaseURL, curDoc)
		if err != nil {
			return err
		}
	}

	return nil
}

func fetchDocument(baseURL string, doc *Document) ([]byte, error) {
	docURL := fmt.Sprintf("%s%s%s", baseURL, doc.Folder, doc.Filename)
	fmt.Printf("fetching %s...\n", docURL)

	resp, err := http.Get(docURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
