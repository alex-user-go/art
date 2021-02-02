# ART
A simple GraphQL server for managing artists and artworks. 
An artist has a name and a set of artworks, each artwork has a title. Artwork can only belong to a single artist.
## Not invented here
This repo was inspired by the great [Kat Zien speech (GopherCon 2018)](https://www.youtube.com/watch?v=oL6JBUk6tj0&t=245s).

GraphQl mutations are based on the [Anemic Mutations](https://xuorig.medium.com/raphql-mutation-design-anemic-mutations-dd107ba70496),
allowing users to update any field independently without optional fields

 
## Examples:
### Adding:
```graphql
mutation {
  addArtwork(title: "44", artistID: 1) {
    id
    title
    artist {
      name
    }
  }
}
mutation {
  addArtist(name: "6655") {
    name
  }
}
```
### Updating:
```graphql
mutation {
  updateArtist(id: 1, actions: { setName: { name: "newname" } }) {
    name
  }
}
mutation {
  updateArtwork(
    id: 1
    actions: { setTitle: { title: "newtitle" }, setArtist: { artistID: 5 } }
  ) {
    id
    title
    artist {
      name
      id
    }
  }
}
```
### Deleting:
```graphql
mutation {
  deleteArtwork(id: 1)
}
mutation {
  deleteArtist(id: 2)
}
```
### Querying:
To query all artists:
```graphql
query {
  filterArtists(name: "") {
    name
    id
  }
}
```
To filter artists:
```graphql
query {
  filterArtists(name: "partOfTheName") {
    name
    id
  }
}
```
Artist's artworks
```graphql
query {
  artist(id: 1) {
    artworks {
      title
      id
    }
  }
}
```