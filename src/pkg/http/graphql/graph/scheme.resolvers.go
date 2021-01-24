package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/alex-user-go/art/pkg/http/graphql/graph/generated"
	"gitlab.com/alex-user-go/art/pkg/http/graphql/graph/model"
)

func (r *mutationResolver) AddArtist(ctx context.Context, name string) (*model.Artist, error) {
	id, err := r.Adder.AddArtist(name)
	if err != nil {
		return nil, err
	}
	arty, err := r.Lister.GetArtist(id)
	if err != nil {
		return nil, err
	}
	artyGraph := model.Artist{ID: arty.ID, Name: arty.Name, Artworks: []*model.Artwork{}}
	return &artyGraph, nil	
}

func (r *mutationResolver) UpdateArtist(ctx context.Context, id int64, actions *model.ArtistUpdateAction) (*model.Artist, error) {
	if actions.SetName != nil {
		_, err := r.Updater.SetArtistName(id, actions.SetName.Name)
		if err != nil {
			return nil, err
		}	
	}
	arty, err := r.Lister.GetArtist(id)
	if err != nil {
		return nil, err
	}
	artyGraph := model.Artist{ID: arty.ID, Name: arty.Name, Artworks: []*model.Artwork{}}
	return &artyGraph, nil	
}

func (r *mutationResolver) DeleteArtist(ctx context.Context, id int64) (*int64, error) {
	_, err := r.Deleter.DeleteArtist(id)
	if err != nil {
		return nil, err
	}
	return &id, err
}

func (r *mutationResolver) AddArtwork(ctx context.Context, title string, artistID int64) (*model.Artwork, error) {
	id, err := r.Adder.AddArtwork(title, artistID)
	if err != nil {
		return nil, err
	}
	arty, err := r.Lister.GetArtist(artistID)
	if err != nil {
		return nil, err
	}
	aw, err := r.Lister.GetArtwork(id)
	if err != nil {
		return nil, err
	}
	artyGraph := model.Artist{ID: arty.ID, Name: arty.Name, Artworks: []*model.Artwork{}}
	awGraph := model.Artwork{ID: id, Title: aw.Title, Artist: &artyGraph}
	return &awGraph, nil
}

func (r *mutationResolver) UpdateArtwork(ctx context.Context, id int64, actions *model.ArtworkUpdateAction) (*model.Artwork, error) {
	if actions.SetTitle != nil {
		_, err := r.Updater.SetArtworkTitle(id, actions.SetTitle.Title)
		if err != nil {
			return nil, err
		}		
	}
	if actions.SetArtist != nil{
		_, err := r.Updater.SetArtworkArtist(id, actions.SetArtist.ArtistID)
		if err != nil {
			return nil, err
		}
	}
	aw, err := r.Lister.GetArtwork(id)
	if err != nil {
		return nil, err
	}
	arty, err := r.Lister.GetArtist(aw.ArtistID)
	if err != nil {
		return nil, err
	}
	artyGraph := model.Artist{ID: arty.ID, Name: arty.Name, Artworks: []*model.Artwork{}}
	awGraph := model.Artwork{ID: id, Title: aw.Title, Artist: &artyGraph}		
	return &awGraph, nil	
}

func (r *mutationResolver) DeleteArtwork(ctx context.Context, id int64) (*int64, error) {
	id, err := r.Deleter.DeleteArtwork(id)
	if err != nil {
		return nil, err
	}
	return &id, err
}

func (r *queryResolver) FilterArtists(ctx context.Context, name string) ([]*model.Artist, error) {
	ars, err := r.Lister.GetArtistByNameFilter(name)
	if err != nil {
		return nil, err
	}
	artistsGraph := make([]*model.Artist, len(ars))
	for i, v := range ars {
		arGraph := model.Artist{ID: v.ID, Name: v.Name}
		artistsGraph[i] = &arGraph
	}
	return artistsGraph, nil
}

func (r *queryResolver) Artist(ctx context.Context, id int64) (*model.Artist, error) {
	arty, err := r.Lister.GetArtist(id)
	if err != nil {
		return nil, err
	}
	aworks, err := r.Lister.GetArtistArtworks(id)
	if err != nil {
		return nil, err
	}
	arworksGraph := make([]*model.Artwork, len(aworks))
	for i, v := range aworks {
		aworkGraph := model.Artwork{ID: v.ID, Title: v.Title, Artist: nil}
		arworksGraph[i] = &aworkGraph
	}
	artyGraph := model.Artist{ID: arty.ID, Name: arty.Name, Artworks: arworksGraph}
	return &artyGraph, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
