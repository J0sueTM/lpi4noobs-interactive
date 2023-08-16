package artifact

type Artifact interface {
	HasCache(string) bool
	LoadCache() error
	WriteCache() error
	FetchCache() error
}
