package listing

//Artist defines a artist to list
type Artist struct {
	ID         int64
	Name       string
	ArtworksID []int64
}

//Artwork defines an Artwork to list
type Artwork struct {
	ID         int64
	Title      string
	ArtistID   int64
}

// Service provides  adding operations.
type Service interface{
	GetArtist(id int64) (*Artist, error)
	GetArtwork(id int64) (*Artwork, error)
	GetArtistArtworks(artistID int64) ([]*Artwork, error)
	GetArtistByNameFilter(filter string)([]*Artist, error)
}

// Repository provides access to the repository.
type Repository interface {
	//GetArtist - get an artist from a repository
	GetArtist(id int64) (*Artist, error)
	//GetArtwork - get an artwork from a repository
	GetArtwork(id int64) (*Artwork, error)
	//GetArtistArtworks - get an artwork slice by artistID
	GetArtistArtworks(artistID int64) ([]*Artwork, error)
	//GetArtistByNameFilter - get an artist slice by name filter
	GetArtistByNameFilter(filter string)([]*Artist, error)
	
	
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetArtist(id int64)  (*Artist, error){
	//any validation can be done here
	return s.r.GetArtist(id)
}

func (s *service) GetArtwork(id int64)  (*Artwork, error){
	//any validation can be done here
	return s.r.GetArtwork(id)
}

func (s *service) GetArtistArtworks(artistID int64)  ([]*Artwork, error){
	//any validation can be done here
	return s.r.GetArtistArtworks(artistID)
}

func (s *service) GetArtistByNameFilter(filter string)  ([]*Artist, error){
	//any validation can be done here
	return s.r.GetArtistByNameFilter(filter)
}


