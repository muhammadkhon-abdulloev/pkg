package logger

type Config struct {
	AddCallerSkip int    `json:"addCallerSkip"`
	Decoder       string `json:"decoder"`
	FilePath      string `json:"filePath"`
	Level         string `json:"level"`
	Production    bool   `json:"development"`
	ToFile        bool   `json:"toFile"`
}
