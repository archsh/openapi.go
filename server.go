package openapi

type Server struct {
    Url         string                    `json:"url"`
    Description string                    `json:"description"`
    Variables   map[string]ServerVariable `json:"variables"`
}

type ServerVariable struct {
    Enum        []string `json:"enum"`
    Default     string   `json:"default"`
    Description string   `json:"description"`
}
