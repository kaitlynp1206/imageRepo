package image

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
)

type ImagesManager struct {
	db     *sql.DB
	bucket *blob.Bucket
}

func NewImagesManager(ctx context.Context, db *sql.DB) *ImagesManager {
	b, err := blob.OpenBucket(ctx, FileBlobStorage)
	if err != nil {
		log.Fatal(err)
	}
	return &ImagesManager{
		db:     db,
		bucket: b,
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
		http.Error(w, BadRequestMsg, http.StatusBadRequest)
	}
}

//GET Image
func (m *ImagesManager) GetImageHandler(w http.ResponseWriter, r *http.Request) {
	var imgReq ImageRequest

	//Parse request body
	err := json.NewDecoder(r.Body).Decode(&imgReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var img *Image
	if imgReq.UserID != 0 && imgReq.FileName != "" {
		id, err := m.GetImageIDbyPath(r.Context(), fmt.Sprintf("%d/%s", imgReq.UserID, imgReq.FileName))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		img, err = m.GetImageMetaDataByID(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else if imgReq.ImageID != 0 {
		img, err = m.GetImageMetaDataByID(r.Context(), imgReq.ImageID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, BadRequestMsg, http.StatusBadRequest)
	}

	resp, _ := json.Marshal(img)

	fmt.Fprintf(w, string(resp))
	return
}

//POST Image
func (m *ImagesManager) AddImageHandler(w http.ResponseWriter, r *http.Request) {
	var imgReq ImageRequest

	//Parse request body
	err := json.NewDecoder(r.Body).Decode(&imgReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Generate path
	path := fmt.Sprintf("%d/%s", imgReq.UserID, imgReq.FileName)

	//Add Image to DB
	err = m.AddImage(r.Context(), imgReq.UserID, imgReq.ImageID, path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Write to blobstore
	err = m.bucket.WriteAll(r.Context(), path, imgReq.Payload, &blob.WriterOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("Successfully added image under %s", path))
	return
}

//DELETE Image
func (m *ImagesManager) DeleteImageHandler(w http.ResponseWriter, r *http.Request) {
	var imgReq ImageRequest

	//Parse request body
	err := json.NewDecoder(r.Body).Decode(&imgReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path := fmt.Sprintf("%d/%s", imgReq.UserID, imgReq.FileName)

	m.bucket.Delete(r.Context(), path)

	err = m.DeleteImage(r.Context(), imgReq.UserID, imgReq.FileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("Successfully removed image under %s", path))
	return
}

func (m *ImagesManager) GetImageMetaDataByID(ctx context.Context, imgID int) (*Image, error) {
	selDB, err := m.db.Query("SELECT image_id, path FROM images WHERE image_id=?", imgID)
	if err != nil {
		return nil, err
	}
	var img Image
	for selDB.Next() {
		err = selDB.Scan(&img.ImageID, &img.Path)
	}
	return &img, nil
}

func (m *ImagesManager) GetImageIDbyPath(ctx context.Context, path string) (int, error) {
	selDB, err := m.db.Query("SELECT image_id FROM images WHERE path=?", path)
	if err != nil {
		return 0, err
	}
	var imgID int
	for selDB.Next() {
		err = selDB.Scan(&imgID)
	}
	return imgID, nil
}

func (m *ImagesManager) AddImage(ctx context.Context, usrID, imgID int, path string) error {
	//TODO: Validate input
	//TODO: Add timestamps for searching/sorting

	insForm, err := m.db.Prepare("INSERT INTO images (path) VALUES(?)")
	if err != nil {
		return err
	}
	insForm.Exec(path)

	//Get new imageID
	imgID, err = m.GetImageIDbyPath(ctx, path)
	if err != nil {
		return err
	}

	//Update linked table
	insForm, err = m.db.Prepare("INSERT INTO users_images (user_id, image_id) VALUES(?,?)")
	if err != nil {
		return err
	}
	insForm.Exec(usrID, imgID)

	//TODO: Return userId
	return nil
}

func (m *ImagesManager) DeleteImage(ctx context.Context, userId int, fileName string) error {
	//TODO: Validate input - ID exists
	path := fmt.Sprintf("%d/%s", userId, fileName)
	delForm, err := m.db.Prepare("DELETE FROM images WHERE path=?")
	if err != nil {
		return err
	}
	delForm.Exec(path)

	id, err := m.GetImageIDbyPath(ctx, path)
	if err != nil {
		return err
	}
	//add del for blobstore

	delForm, err = m.db.Prepare("DELETE FROM users_images WHERE image_id=?")
	if err != nil {
		return err
	}
	delForm.Exec(id)

	//TODO: Validate change
	return nil
}
