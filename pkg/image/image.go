package image

type Image struct {
	ImageID   int    `json:"image_id"`
	Path      string `json:"path"`
	ImageName string `json:"image_name"`
	Payload   []byte `json:"payload"`
}

type ImageRequest struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	ImageID  int    `json:"image_id"`
	FileName string `json:"file_name"`
	Payload  []byte `json:"payload"`
}
