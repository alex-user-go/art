package adding

import (
	"testing"

	"gitlab.com/alex-user-go/art/pkg/listing"
)
func TestAddArtist(t *testing.T){
	artistName := "nino"

	mR := new(mockStorage)

	s := NewService(mR)
	_, err := s.AddArtist(artistName)
	if err != nil{
		t.Fatal("\t\tShould be able to add a new artist. ",err)
	}

	artowrkTitle := "the table"
	_, err = s.AddArtwork(artowrkTitle, 1)
	if err != nil{
		t.Fatal("\t\tShould be able to add a new artwork. ",err)
	}
}

type mockStorage struct {
	artists []listing.Artist
	artworks []listing.Artwork
}

func (m * mockStorage)AddArtist(name string)  (int64, error){
	artist:= listing.Artist{Name: name}
	m.artists = append(m.artists,artist )
	return 1, nil
}

func (m * mockStorage) AddArtwork(title string, artistID int64) (int64, error){
	artwork := listing.Artwork{Title: title, ArtistID: artistID}
	m.artworks = append(m.artworks,artwork )
	return 1, nil
}