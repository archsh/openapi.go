package openapi

type Tag struct {
    Name        string `json:"name"`
    Description string `json:"description"`
}

type Reference struct {
    Ref string `json:"$ref"`
}
