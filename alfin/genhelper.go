package alfin

import "go.uber.org/zap"

type GenHelper struct {
	Logger *zap.Logger
}

func NewGenHelper(logger *zap.Logger) *GenHelper {
	return &GenHelper{Logger: logger}
}

func DevGenHelper() *GenHelper {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &GenHelper{Logger: logger}
}
