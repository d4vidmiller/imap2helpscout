package helpscout

import "strconv"

type reqUploadAttachment struct {
	Name     string `json:"fileName"`
	MimeType string `json:"mimeType"`
	Data     []byte `json:"data"`
}

// UploadAttachment uploads an attachment to the given conversation > thread
func (h *HelpScout) UploadAttachment(conversationID uint64, threadID uint64, name string, mimeType string, data []byte) (resp []byte, err error) {
	_, _, resp, err = h.Exec(
		"conversations/"+strconv.FormatUint(conversationID, 10)+"/threads/"+strconv.FormatUint(threadID, 10)+"/attachments",
		reqUploadAttachment{
			Name:     name,
			MimeType: mimeType,
			Data:     data,
		},
		nil,
		"",
	)
	return
}
