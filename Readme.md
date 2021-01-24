
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