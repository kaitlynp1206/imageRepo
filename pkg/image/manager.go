package image

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type ImagesManager struct {
	db *sql.DB
	//bucket *storage.Bucket
}

func NewImagesManager(ctx context.Context, db *sql.DB) *ImagesManager {
	return &ImagesManager{
		db: db,
		//blobStore: storage.NewBucket(ctx),
	}
}

//Image Handler
func (m *ImagesManager) ImageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		m.GetImageHandler(w, r)
	case "POST":
		m.AddImageHandler(w, r)
	case "DELETE":
		m.DeleteImageHandler(w, r)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

//GET Images
func (m *ImagesManager) GetImageHandler(w http.ResponseWriter, r *http.Request) {
	var imgReq ImageRequest

	//Parse request body
	err := json.NewDecoder(r.Body).Decode(&imgReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("Retrieved Image %s", imgReq.ImageID))
	return
}

//POST to Images
func (m *ImagesManager) AddImageHandler(w http.ResponseWriter, r *http.Request) {
	var imgReq ImageRequest

	//Parse request body
	err := json.NewDecoder(r.Body).Decode(&imgReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Generate blob
	path := fmt.Sprintf("%d/%s", imgReq.UserID, imgReq.Username)
	/*err := m.blobStore.WriteAll(ctx, path, imgReq.Payload, blob.WriterOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}*/

	//Add Image to DB
	err = m.AddImage(r.Context(), imgReq.UserID, imgReq.ImageID, path, imgReq.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("Successfully added image as %s", path))
	return
}

//DELETE Image
func (m *ImagesManager) DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	var imgReq Image

	//Parse request body
	err := json.NewDecoder(r.Body).Decode(&imgReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = m.DeleteImage(r.Context(), imgReq.ImageID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("Successfully removed image %s", imgReq.ImageID))
	return
}

func (m *ImagesManager) GetImageMetaDataByID(ctx context.Context, imgID int) (*Image, error) {
	selDB, err := m.db.Query("SELECT image_id, path, description WHERE image_id=?", imgID)
	if err != nil {
		return nil, err
	}
	var img Image
	for selDB.Next() {
		err = selDB.Scan(&img.ImageID, &img.Path, &img.Description)
	}
	return &img, nil
}

func (m *ImagesManager) GetImageIDByPath(ctx context.Context, path string) (int, error) {
	selDB, err := m.db.Query("SELECT image_id, path, description WHERE path=?", path)
	if err != nil {
		return 0, err
	}
	var imgID int
	for selDB.Next() {
		err = selDB.Scan(&imgID)
	}
	return imgID, nil
}

func (m *ImagesManager) AddImage(ctx context.Context, usrID, imgID int, path, desc string) error {
	//TODO: Validate input
	insForm, err := m.db.Prepare("INSERT INTO images(path, description) VALUES(?,?)")
	if err != nil {
		return err
	}
	insForm.Exec(path, desc)

	//Get new imageID
	imgID, err = m.GetImageIDByPath(ctx, path)
	if err != nil {
		return err
	}

	//Update linked table
	insForm, err = m.db.Prepare("INSERT INTO users_images(user_id, image_id) VALUES(?,?)")
	if err != nil {
		return err
	}
	insForm.Exec(usrID, imgID)

	//TODO: Return userId
	return nil
}

func (m *ImagesManager) DeleteImage(ctx context.Context, imgID int) error {
	//TODO: Validate input - ID exists
	delForm, err := m.db.Prepare("DELETE FROM images WHERE image_id=?")
	if err != nil {
		return err
	}
	delForm.Exec(imgID)

	delForm, err = m.db.Prepare("DELETE FROM users_images WHERE image_id=?")
	if err != nil {
		return err
	}
	delForm.Exec(imgID)

	//TODO: Validate change
	return nil
}
