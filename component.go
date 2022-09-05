package openapi

type Component struct {
    Schemas         map[string]Schema         `json:"schemas"`
    Responses       map[string]Response       `json:"responses"`
    Parameters      map[string]Parameter      `json:"parameters"`
    Examples        map[string]Example        `json:"examples"`
    RequestBodies   map[string]RequestBody    `json:"requestBodies"`
    Headers         map[string]Header         `json:"headers"`
    SecuritySchemes map[string]SecurityScheme `json:"securitySchemes"`
    Links           map[string]Link           `json:"links"`
    Callbacks       map[string]Callback       `json:"callbacks"`
}

type Schema struct {
}

type Response struct {
    Description string               `json:"description"`
    Headers     map[string]Header    `json:"headers"`
    Content     map[string]MediaType `json:"content"`
    Links       map[string]Link      `json:"links"`
}

type Parameter struct{}

type Example struct {
    Summary       string      `json:"summary"`
    Description   string      `json:"description"`
    Value         interface{} `json:"value"`
    ExternalValue string      `json:"externalValue"`
}

type RequestBody struct {
    Description string               `json:"description"`
    Content     map[string]MediaType `json:"content"`
    Required    bool                 `json:"required"`
}

type MediaType struct {
    Schema   Schema              `json:"schema"`
    Example  interface{}         `json:"example"`
    Examples map[string]Example  `json:"examples"`
    Encoding map[string]Encoding `json:"encoding"`
}

type Encoding struct{}

type Header struct{}

type SecurityScheme struct{}

type Link struct {
    OperationRef string                 `json:"operationRef"`
    OperationId  string                 `json:"operationId"`
    Parameters   map[string]interface{} `json:"parameters"`
    RequestBody  interface{}            `json:"requestBody"`
    Description  string                 `json:"description"`
    Server       Server                 `json:"server"`
}

type Callback struct{}
