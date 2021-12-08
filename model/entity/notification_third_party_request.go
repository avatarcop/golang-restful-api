package entity

type Headings struct {
	En string `json:"en"`
}

type Contents struct {
	En string `json:"en"`
}

type OnesignalPayload struct {
	AppId                     string   `json:"app_id"`
	IncludeExternalUserIds    []string `json:"include_external_user_ids"`
	ChannelForExternalUserIds string   `json:"channel_for_external_user_ids"`
	Headings                  Headings `json:"headings"`
	Contents                  Contents `json:"contents"`
}
