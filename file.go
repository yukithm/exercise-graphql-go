package exercise_graphql_go

type File struct {
	Path    string  `json:"path"`
	Size    int64   `json:"size"`
	IsDir   bool    `json:"isDir"`
	Entries []*File `json:"entries"`
}
