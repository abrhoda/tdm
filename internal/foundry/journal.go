package foundry

import "encoding/json"

type Journal struct {
	ID    string `json:"_id"`
	Pages []Page `json:"pages"`
}

type Page struct {
	ID   string
	Name string
	Text string
}

func (p *Page) UnmarshalJSON(b []byte) error {
	var temp struct {
		ID   string         `json:"_id"`
		Name string         `json:"name"`
		Text map[string]any `json:"text"`
	}

	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	p.ID = temp.ID
	p.Name = temp.Name
	p.Text = temp.Text["content"].(string)

	return nil
}
