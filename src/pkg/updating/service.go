package updating

// Service provides  adding operations.
type Service interface{
	SetArtistName(id int64, name string) (int64, error)	
	SetArtworkTitle(id int64, title string) (int64, error)
	SetArtworkArtist(id int64, artistID int64) (int64, error)
}

// Repository provides access to the repository.
type Repository interface {
	SetArtistName(id int64, name string) (int64, error)
	SetArtworkTitle(id int64, title string) (int64, error)
	SetArtworkArtist(id int64, artistID int64) (int64, error)
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) SetArtistName(id int64, name string) (int64, error){
	//any validation can be done here
	return s.r.SetArtistName(id, name)
}
func (s *service)	SetArtworkTitle(id int64, title string) (int64, error){
	//any validation can be done here
	return s.r.SetArtworkTitle(id, title)
}
func (s *service) SetArtworkArtist(id int64, artistID int64)  (int64, error){
	//any validation can be done here
	return s.r.SetArtworkArtist(id, artistID)
}


