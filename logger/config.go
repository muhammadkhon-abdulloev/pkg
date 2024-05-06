package logger

type Config struct {
	AddCallerSkip int    `json:"addCallerSkip" mapstructure:"LOGGER_ADD_CALLER_SKIP"`
	Decoder       string `json:"decoder" mapstructure:"LOGGER_DECODER"`
	FilePath      string `json:"filePath" mapstructure:"LOGGER_FILE_PATH"`
	Level         string `json:"level" mapstructure:"LOGGER_LEVEL"`
	Production    bool   `json:"development" mapstructure:"LOGGER_PRODUCTION"`
	ToFile        bool   `json:"toFile" mapstructure:"LOGGER_TO_FILE"`
}
