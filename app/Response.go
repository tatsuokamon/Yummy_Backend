package backend

import (
	"local.package.ytbdtc"
)

type Response struct {
	Items []ytbdtc.Item `json:"items"`
	NextPageToken string `json:"nextPageToken"`
	Success bool `json:"success"`
	Message string `json:"message"`
}
