package domain

type MenuButton struct {
	Type      string        `json:"Type,omitempty"`
	Name      string        `json:"Name,omitempty"`
	Key       string        `json:"Key,omitempty"`
	Url       string        `json:"Url,omitempty"`
	MediaId   string        `json:"MediaId,omitempty"`
	AppId     string        `json:"AppId,omitempty"`
	PagePath  string        `json:"PagePath,omitempty"`
	SubButton []*MenuButton `json:"SubButton,omitempty"`
}
