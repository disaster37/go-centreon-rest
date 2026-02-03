package api

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"emperror.dev/errors"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// MediaService defines the interface for managing media in Centreon.
type MediaService interface {

	// Create creates a new media.
	Create(media *MediaCreateRequest) (mediaResponse *MediaCreateResponse, err error)

	// Update updates an existing media by its ID.
	Update(id int64, name string, contend []byte) (mediaResponse *MediaCreateOrUpdateResponseResult, err error)

	// Get retrieves a media by its ID.
	// Need to have PR https://github.com/centreon/centreon/pull/9329
	Get(id int64) (mediaResponse *MediaReponse, err error)

	// GetByName retrieves a media by its Name.
	// It use Find method to get the media
	GetByName(name string) (mediaResponse *MediaListReponse, err error)

	// Find retrieves a list of media based on the provided options.
	Find(options *ListOptions) (mediaListResponse *ListResponse[MediaListReponse], err error)

	// Delete deletes a media by its ID.
	// Need to have PR https://github.com/centreon/centreon/pull/9329
	Delete(id int64) (err error)
}

// DefaultMediaService implements MediaService
type DefaultMediaService struct {
	client *resty.Client
	logger *logrus.Entry
}

// NewMediaService creates a new instance of DefaultMediaService
func NewMediaService(client *resty.Client, logger *logrus.Entry) MediaService {
	return &DefaultMediaService{
		client: client,
		logger: logger.WithField("service", "media"),
	}
}

// Create creates a new media
func (m *DefaultMediaService) Create(media *MediaCreateRequest) (mediaResponse *MediaCreateResponse, err error) {

	m.logger.Debugf("Create Media: %+v", media)

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(media); err != nil {
		return nil, errors.Wrap(err, "validation error on create media")
	}

	mediaResponse = new(MediaCreateResponse)

	formData := map[string]string{
		"directory": media.Directory,
	}

	response, err := m.client.R().
		SetFileReader("data", media.Name, bytes.NewReader(media.Data)).
		SetMultipartFormData(formData).
		SetResult(mediaResponse).
		Post("/configuration/medias")

	if err != nil {
		return nil, errors.Wrap(err, "error during create media request")
	}

	if response.IsError() {

		return nil, errors.Errorf("create media failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	m.logger.Debugf("Media created successfully: %+v", mediaResponse)
	return mediaResponse, nil
}

// Update updates an existing media
func (m *DefaultMediaService) Update(id int64, name string, contend []byte) (responseMedia *MediaCreateOrUpdateResponseResult, err error) {

	m.logger.Debugf("Update Media with id: %d", id)

	if id <= 0 {
		return nil, errors.Errorf("invalid media id: %d", id)
	}

	responseMedia = new(MediaCreateOrUpdateResponseResult)

	response, err := m.client.R().
		SetFileReader("data", name, bytes.NewReader(contend)).
		SetPathParam("mediaId", fmt.Sprintf("%d", id)).
		SetResult(responseMedia).
		Post("/configuration/medias/{mediaId}/content")

	if err != nil {
		return nil, errors.Wrapf(err, "error during update media request: %s", response.String())
	}
	if response.IsError() {
		return nil, errors.Errorf("update media failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	m.logger.Debugf("Media with id %d updated successfully", id)
	return responseMedia, nil
}

// List retrieves a list of media based on the provided options.
func (m *DefaultMediaService) Find(options *ListOptions) (mediaListResponse *ListResponse[MediaListReponse], err error) {
	if options == nil {
		options = &ListOptions{}
	}

	m.logger.Debugf("Find media with options: %+v", options)

	mediaListResponse = new(ListResponse[MediaListReponse])

	response, err := m.client.R().
		SetQueryParams(options.GetQueryParams()).
		SetResult(mediaListResponse).
		Get("/configuration/medias")

	m.logger.Debugf("Response from find media: %s", response.String())

	if err != nil {
		return nil, errors.Wrap(err, "error during find media request")
	}

	if response.IsError() {
		return nil, errors.Errorf("list media failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return mediaListResponse, nil
}

// Delete deletes a media by its ID
func (m *DefaultMediaService) Delete(id int64) (err error) {
	m.logger.Debugf("Delete media with Id: %d", id)

	if id <= 0 {
		return errors.Errorf("invalid media id: %d", id)
	}

	response, err := m.client.R().
		SetPathParam("mediaId", fmt.Sprintf("%d", id)).
		Delete("/configuration/medias/{mediaId}")

	m.logger.Debugf("Response from delete media: %s", response.String())

	if err != nil {
		return errors.Wrapf(err, "error during delete media request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil
		}
		return errors.Errorf("delete media failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	m.logger.Debugf("Media with id %d deleted successfully", id)
	return nil
}

// Get retrieves a media by its ID.
func (m *DefaultMediaService) Get(id int64) (mediaResponse *MediaReponse, err error) {
	m.logger.Debugf("Get media with Id: %d", id)

	if id <= 0 {
		return nil, errors.Errorf("invalid media id: %d", id)
	}

	mediaResponse = new(MediaReponse)

	response, err := m.client.R().
		SetPathParam("mediaId", fmt.Sprintf("%d", id)).
		SetResult(mediaResponse).
		Get("/configuration/medias/{mediaId}")

	m.logger.Debugf("Response from get media: %s", response.String())

	if err != nil {
		return nil, errors.Wrapf(err, "error during get media request")
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusNotFound {
			return nil, nil
		}
		return nil, errors.Errorf("get media failed with status code %d and message %s", response.StatusCode(), response.String())
	}

	return mediaResponse, nil
}

// GetByName retrieves a media by its Name.
func (m *DefaultMediaService) GetByName(name string) (mediaResponse *MediaListReponse, err error) {
	m.logger.Debugf("Get media with Name: %s", name)

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("media name cannot be empty")
	}

	mediaListResponse, err := m.Find(&ListOptions{
		Search: map[string]interface{}{
			"name": name,
		},
	})

	if err != nil {
		return nil, errors.Wrapf(err, "error during find media by Name: %s", name)
	}

	if mediaListResponse.Meta.Total == 1 {
		return &mediaListResponse.Result[0], nil
	}

	return nil, nil
}
