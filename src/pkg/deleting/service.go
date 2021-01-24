package deleting

// Service provides  deleting operations.
type Service interface{
	DeleteArtist(id int64) (int64, error)
	DeleteArtwork(id int64) (int64, error)
}

// Repository provides access to the repository.
type Repository interface {
	//delete an artist to a repository
	DeleteArtist(id int64) (int64, error)
	//delete an artwork to a repository
	DeleteArtwork(id int64) (int64, error)
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) 	DeleteArtist(id int64) (int64, error){
	//any validation can be done here
	return s.r.DeleteArtist(id)
}
func (s *service) DeleteArtwork(id int64) (int64, error) {
	//any validation can be done here
	return s.r.DeleteArtwork(id)
}


