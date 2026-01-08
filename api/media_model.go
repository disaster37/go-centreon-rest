package api

// MediaCreateRequest represents the payload to create a new media in Centreon.
type MediaCreateRequest struct {
	Name 	string `json:"name,omitempty" validate:"required"`
	Directory string `json:"directory" validate:"required"`
	Data      []byte `json:"data" validate:"required"`
}

// MediaCreateResponse represents the response after creating media in Centreon.
type MediaCreateResponse struct {
	Result []MediaCreateResponseResult `json:"result"`
	Errors []MediaCreateResponseError  `json:"errors,omitempty"`
}

// MediaCreateResponseResult represents a successful media creation result.
type MediaCreateResponseResult struct {
	Id        int64  `json:"id"`
	Filename  string `json:"filename"`
	Directory string `json:"directory"`
	Md5       string `json:"md5"`
}

// MediaCreateResponseError represents an error that occurred during media creation.
type MediaCreateResponseError struct {
	Filename  string `json:"filename"`
	Directory string `json:"directory"`
	Reason    string `json:"reason"`
}

// MediaReponse represents a media in Centreon.
type MediaReponse struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Directory string `json:"directory"`
	Url       string `json:"url"`
	Md5       string `json:"md5"`
}
