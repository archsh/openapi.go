package openapi

type Operation struct {
    Tags         []string            `json:"tags"`
    Summary      string              `json:"summary"`
    Description  string              `json:"description"`
    ExternalDocs Documentation       `json:"externalDocs"`
    OperationId  string              `json:"operationId"`
    Parameters   []Parameter         `json:"parameters"`
    RequestBody  RequestBody         `json:"requestBody"`
    Responses    Response            `json:"responses"`
    Callbacks    map[string]Callback `json:"callbacks"`
    Deprecated   bool                `json:"deprecated"`
    Security     []Security          `json:"security"`
    Servers      []Server            `json:"servers"`
}
