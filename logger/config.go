package logger

type Config struct {
	AddCallerSkip int    `json:"addCallerSkip" env:"LOGGER_ADD_CALLER_SKIP"`
	Decoder       string `json:"decoder" env:"LOGGER_DECODER"`
	FilePath      string `json:"filePath" env:"LOGGER_FILE_PATH"`
	Level         string `json:"level" env:"LOGGER_LEVEL"`
	Production    bool   `json:"development" env:"LOGGER_PRODUCTION"`
	ToFile        bool   `json:"toFile" env:"LOGGER_TO_FILE"`
}
