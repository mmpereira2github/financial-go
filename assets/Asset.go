package assets

// Asset represent an investement
type Asset interface {
	ID() int
	Name() string
}

type asset struct {
	id   int
	name string
}

// NewAsset creates a new asset
func NewAsset(id int, name string) Asset { return &asset{id, name} }

func (a *asset) ID() int      { return a.id }
func (a *asset) Name() string { return a.name }
