package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

// logger 对象
func init() {
	test()
}

// 生产定制化的日志记录器
func test() {

	//定制json 格式的日志核心
	//编码器(如何写入日志)     这些都可以定制
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	file, _ := os.Create("./test.log") //
	//日志输出的位置
	writeSyncer := zapcore.AddSync(file)

	//日志级别 可以从配置文件中 自定义什么级别
	// viper 配置日志
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)
	// 自定义的日志要添加zap.AddCaller() 来记录日志调用的位置
	logger := zap.New(core, zap.AddCaller())

	var url string = "https://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	logger.Info("Success..",
		zap.String("statusCode", resp.Status),
		zap.String("url", url))

}

func test1() {
	// zap 默认有两种日志记录器 一个是 Sugared Logger  一个是Logger
	// 输出都是json格式 再标准输出到终端

	// 这个logger会记录 这个日志调用的位置
	logger, err := zap.NewProduction()
	SugaredLogger := logger.Sugar()
	if err != nil {
		panic(err)
	}
	//
	resp, err := http.Get("http://www.baidu.com")
	// string fild 级别
	// 用这个比较的好
	logger.Info("Success..",
		zap.String("statusCode", resp.Status),
		zap.String("url", "http://www.baidu.com"))

	//这个输出要格式化
	SugaredLogger.Info("Success..", resp.Status, "http://www.baidu.com")
	resp.Body.Close()
}
