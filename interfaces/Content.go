package interfaces

type CONTENT struct{
    TYPE        string `json:"type" binding:"required" validate:"omitempty,oneof:link,extension"`
    LINK        string `json:"link" validate:"omitempty,required_unless=TYPE link"`
    HIGHLIGHTER string `json:"highlighter" binding:"required"`
    CONTEXT     string `json:"context"`
    TEXT        string `json:"text"`
}