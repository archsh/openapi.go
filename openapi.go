package openapi

type OpenAPI struct {
    Openapi      string        `json:"openapi"`
    Info         Info          `json:"info"`
    Servers      []Server      `json:"servers"`
    Paths        []Path        `json:"paths"`
    Components   []Component   `json:"components"`
    Security     []Security    `json:"security"`
    Tags         []Tag         `json:"tag"`
    ExternalDocs Documentation `json:"externalDocs"`
}

type Documentation struct {
    Description string `json:"description"`
    Url         string `json:"url"`
}
