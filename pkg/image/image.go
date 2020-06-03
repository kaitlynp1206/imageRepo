package image

type Image struct {
	ImageID   int
	Path      string
	ImageName string
	Payload   []byte
}

type ImageRequest struct {
	UserID   int `json:"user_id"`
	Username string
	ImageID  int
	FileName string `json:"file_name"`
	Payload  []byte
}
