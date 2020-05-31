package image

import (
	"context"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/<driver>"
)

type ImagesManager struct {
	db *sql.DB
	blobStore *blob.Bucket
}

func NewUsers(db *sql.DB) *ImagesManager {
	return &ImagesManager{
		db: db,
	}
}

func (m *ImagesManager) AddHandler(w http.ResponseWriter, r *http.Request) {

}

func (m *ImagesManager) DeleteHandler(w http.ResponseWriter, r *http.Request) {

}

func (m *ImagesManager) AddImage(ctx context.Context, usrID int, payload []byte, name, desc string) error {
	//TODO: Validate input
	path := fmt.Sprintf("%d/%s", userId, name)
	err := m.blobStore.WriteAll(ctx, path, payload, blob.WriterOptions{})
	if err != nil {
		return err
	}
	insForm, err := m.db.Prepare("INSERT INTO images(path, description) VALUES(?,?)")
	if err != nil {
		return err
	}
	insForm.Exec(path, desc)
	
	//Get new imageID
	imgID = GetImageIDByPath(ctx, path)

	//Update linked table
	insForm, err := m.db.Prepare("INSERT INTO users_images(user_id, image_id) VALUES(?,?)")
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

	delForm, err := m.db.Prepare("DELETE FROM users_images WHERE image_id=?")
	if err != nil {
		return err
	}
	delForm.Exec(imgID)

	//TODO: Validate change
	return nil
}

func (m *ImagesManager) GetImageMetaDataByID(ctx context.Context, imgID int) (Image, error) {
	selDB, err := m.db.Query("SELECT image_id, path, description WHERE image_id=?", imgID)
	if err != nil {
		return nil
	}
	var img Image
	for selDB.Next() {
		err = selDB.Scan(&img.ImageID, &img.Path, &img.Description)
	}
	return img, nil
}

func (m *ImagesManager) GetImageIDByPath(ctx context.Context, path string) (int, error) {
	selDB, err := m.db.Query("SELECT image_id, path, description WHERE path=?", path)
	if err != nil {
		return nil
	}
	var imgID int
	for selDB.Next() {
		err = selDB.Scan(&imgID)
	}
	return imgID, nil
}
