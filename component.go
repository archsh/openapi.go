package openapi

type Component struct {
    Schemas         map[string]Schema         `json:"schemas"`
    Responses       map[string]Response       `json:"responses"`
    Parameters      map[string]Parameter      `json:"parameters"`
    Examples        map[string]Example        `json:"examples"`
    RequestBodies   map[string]RequestBody    `json:"requestBodies"`
    Headers         map[string]Header         `json:"headers"`
    SecuritySchemes map[string]SecurityScheme `json:"SecuritySchemes"`
    Links           map[string]Link           `json:"links"`
    Callbacks       map[string]Callback       `json:"callbacks"`
}

type Schema struct {
}

type Response struct{}

type Parameter struct{}

type Example struct{}

type RequestBody struct{}

type Header struct{}

type SecurityScheme struct{}

type Link struct{}

type Callback struct{}
