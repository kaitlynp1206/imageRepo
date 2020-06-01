package image

type Image struct {
	ImageID     int
	Path        string
	Description string
	Payload     []byte
}

type ImageRequest struct {
	UserID      int
	Username    string
	ImageID     int
	Path        string
	Description string
	Payload     []byte
}
