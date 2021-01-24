package sqlite

import (
	"database/sql"

	//sqlite realization
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/alex-user-go/art/pkg/listing"
)

//Storage - sqlite db
type Storage struct{
	db *sql.DB
}

//NewStorage - create and init new sqlite storage
func NewStorage()(*Storage, error){
	db, err := sql.Open("sqlite3", "./art.db")
	if err != nil {
		return nil, err	
	}
	s:=  &Storage{db}

	if err := s.createTables(); err !=nil{
		return nil, err
	}
	return s, nil
}

func (s *Storage) createTables() error{
	
	if err := s.createTableArtist(); err !=nil{
		return err
	}
	if err := s.createTableArtwork(); err !=nil{
		return err
	}
	if err := s.setConstraits(); err != nil{
		return err
	}
	return nil
}

func (s *Storage) setConstraits() error{
	_, err := s.db.Exec("PRAGMA foreign_keys = ON", nil)
	if err != nil {
		return err
	}
	return nil
}
func (s *Storage) createTableArtist() error{
	sqlStmt := `create table if not exists
				Artist (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE)`
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
} 
func (s *Storage) createTableArtwork() error{
	sqlStmt := `create table if not exists
	Artwork (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, artistID INTEGER NOT NULL, FOREIGN KEY(artistID) REFERENCES Artist(id))`
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
} 
//AddArtist add a new artist to the db
func (s *Storage) AddArtist(name string) (int64, error){
	stmt, err := s.db.Prepare("insert into Artist(name) values(?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(name)
	if err != nil {
		return 0, err
	}
	lid, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lid, nil	
}
//AddArtwork add a new artwork to the db
func (s *Storage) AddArtwork(title string, artistID int64) (int64, error){
	stmt, err := s.db.Prepare("insert into Artwork(title, artistID) values(?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(title, artistID)
	if err != nil {
		return 0, err
	}
	lid, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lid, nil	
}

//GetArtist - get artist by id
func (s *Storage) GetArtist(id int64)  (*listing.Artist, error){
	sqlStmt := `SELECT id, name FROM Artist WHERE id = ?`
	var ar listing.Artist
	err := s.db.QueryRow(sqlStmt, id).
						Scan(&ar.ID, &ar.Name)
	if err != nil{
		return nil, err
	}
	return &ar, nil
}
//GetArtwork - get an artwork by id
func (s *Storage) GetArtwork(id int64)  (*listing.Artwork, error){
	sqlStmt := `SELECT id, title, artistID FROM Artwork WHERE id = ?`
	var arwork listing.Artwork
	err := s.db.QueryRow(sqlStmt, id).
						Scan(&arwork.ID, &arwork.Title, &arwork.ArtistID)
	if err != nil{
		return nil, err
	}
	return &arwork, nil
}


//DeleteArtist - delete artist by id
func (s *Storage) DeleteArtist(id int64)  (int64, error){
	sqlStmt := `delete FROM Artist WHERE id = ?`
	_, err := s.db.Exec(sqlStmt, id)
	if err != nil{
		return 0, err
	}	
	return id, nil
}
//DeleteArtwork - delete artist by id
func (s *Storage) DeleteArtwork(id int64)  (int64, error){
	sqlStmt := `delete FROM Artwork WHERE id = ?`
	_, err := s.db.Exec(sqlStmt, id)
	if err != nil{
		return 0, err
	}	
	return id, nil
}



//GetArtistArtworks - get artworks by artistID
func (s *Storage) GetArtistArtworks(artistID int64) ([]*listing.Artwork, error){
	sqlStmt := `SELECT id, title, artistID FROM Artwork WHERE artistID = ?`
	rows, err := s.db.Query(sqlStmt,  artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var arworks []*listing.Artwork
	for rows.Next() {
		var arwork listing.Artwork		
		err = rows.Scan(&arwork.ID, &arwork.Title, &arwork.ArtistID)
		if err != nil {
			return nil, err
		}	
		arworks = append(arworks, &arwork)		
	}
	return arworks, nil
}


//GetArtistByNameFilter - get atrist name filter
func (s *Storage) GetArtistByNameFilter(filter string) ([]*listing.Artist, error){
	sqlStmt := `SELECT id, name FROM Artist WHERE  instr(name, ?)>0`
	rows, err := s.db.Query(sqlStmt,  filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ars []*listing.Artist
	for rows.Next() {
		var ar listing.Artist		
		err = rows.Scan(&ar.ID, &ar.Name)
		if err != nil {
			return nil, err
		}	
		ars = append(ars, &ar)		
	}
	return ars, nil
}
//SetArtistName - set artist name by id
func (s *Storage) SetArtistName(id int64, name string) (int64, error){
	stmt, err := s.db.Prepare("update Artist set name = ? where id = ?")
	if err != nil {
		return  0, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, id)
	if err != nil {
		return  0, err
	}
	return id, nil
}
//SetArtworkTitle - set artwork name by id
func (s *Storage)  SetArtworkTitle(id int64, title string) (int64, error){
	stmt, err := s.db.Prepare("update Artwork set title = ? where id = ?")
	if err != nil {
		return  0, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(title, id)
	if err != nil {
		return  0, err
	}
	return id, nil
}
//SetArtworkArtist - set artwork author
func (s *Storage)  SetArtworkArtist(id int64, artistID int64)  (int64, error){
	stmt, err := s.db.Prepare("update Artwork set artistID = ? where id = ?")
	if err != nil {
		return  0, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(artistID, id)
	if err != nil {
		return  0, err
	}
	return id, nil
}