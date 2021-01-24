package adding

// Service provides  adding operations.
type Service interface{
	AddArtist(name string) (int64, error)
	AddArtwork(title string, artistID int64) (int64, error)
}

// Repository provides access to the repository.
type Repository interface {
	//save an artist to a repository
	AddArtist(name string) (int64, error)
	//save an artwork to a repository
	AddArtwork(title string, artistID int64) (int64, error)
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddArtist(name string)  (int64, error){
	//any validation can be done here
	return s.r.AddArtist(name)
}
func (s *service) AddArtwork(title string, artistID int64) (int64, error){
	//any validation can be done here
	return s.r.AddArtwork(title, artistID)
}


