package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger       *zap.Logger
	loggerWriter *lumberjack.Logger
	location     *time.Location
)

type Config struct {
	Level    string `json:"level"`
	RootDir  string `json:"root_dir"`
	Filename string `json:"filename"`

	// 滚动日志配置
	MaxSize    int    `json:"max_size"`    // 每个日志文件的最大尺寸（以MB为单位）
	MaxBackups int    `json:"max_backups"` // 保留的旧日志文件的最大备份数
	MaxAge     int    `json:"max_age"`     // 最大的日志文件保留天数
	Compress   bool   `json:"compress"`    // 是否压缩旧日志文件
	Env        string `json:"env"`
	TZ         string `json:"tz"`
}

func NewLogger(conf *Config) *zap.Logger {
	var (
		level   zapcore.Level // zap 日志等级
		options []zap.Option  // zap 配置项
	)
	rootPath := RootPath()
	logFileDir := conf.RootDir
	if !filepath.IsAbs(logFileDir) {
		logFileDir = filepath.Join(rootPath, logFileDir)
	}
	if ok, _ := Exists(logFileDir); !ok {
		_ = os.Mkdir(conf.RootDir, os.ModePerm)
	}
	switch conf.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddCaller())
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
		options = append(options, zap.AddCaller())
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		if location == nil {
			encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
		} else {
			encoder.AppendString(time.In(location).Format("2006-01-02 15:04:05.000"))
		}
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(conf.Env + "." + l.String())
	}

	loggerWriter = &lumberjack.Logger{
		Filename:   filepath.Join(logFileDir, conf.Filename),
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
		Compress:   conf.Compress,
	}

	now := time.Now()
	loc, err := time.LoadLocation(conf.TZ)
	if err == nil {
		location = loc
	}
	date := now.In(location).Format(time.DateOnly)
	hook := func(e zapcore.Entry) error {
		if currDate := e.Time.In(location).Format(time.DateOnly); currDate != date {
			date = currDate
			_ = loggerWriter.Rotate()
		}
		return nil
	}
	options = append(options, zap.Hooks(hook))
	logger = zap.New(zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.AddSync(loggerWriter), level), options...)
	return logger
}

// RootPath 获取项目根目录绝对路径
func RootPath() string {
	var rootDir string

	exePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	rootDir = exePath

	tmpDir := os.TempDir()
	if strings.Contains(exePath, tmpDir) {
		_, filename, _, ok := runtime.Caller(0)
		if ok {
			rootDir = filepath.Dir(filepath.Dir(filepath.Dir(filename)))
		}
	}

	return rootDir
}

// Exists 路径是否存在
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
