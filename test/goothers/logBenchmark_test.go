package goothers

import (
	"godemo/internal/goothers/logrusdemo"
	"godemo/internal/goothers/zapdemo"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type tmpData struct {
	Message string
	Code    int
	Err     string
	Url     string
}

func BenchmarkZapLog(b *testing.B) {
	zapLogger := zapdemo.NewZapLogger()

	infoData := &tmpData{}
	errorData := &tmpData{}

	err := faker.FakeData(infoData)
	if err != nil {
		b.Fatal(err)
	}

	err = faker.FakeData(errorData)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zapLogger.Info(infoData.Message, zap.Int("Code", infoData.Code),
			zap.String("Err", infoData.Err), zap.String("Url", infoData.Url))

		zapLogger.Error(errorData.Message, zap.Int("Code", errorData.Code),
			zap.String("Err", errorData.Err), zap.String("Url", errorData.Url))
	}

}

func BenchmarkLogrusLog(b *testing.B) {

	logrusLogger := logrusdemo.NewLogrusLogger()

	infoData := &tmpData{}
	errorData := &tmpData{}

	err := faker.FakeData(infoData)
	if err != nil {
		b.Fatal(err)
	}

	err = faker.FakeData(errorData)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logrusLogger.WithFields(logrus.Fields{
			"Code": infoData.Code,
			"Err":  infoData.Err,
			"Url":  infoData.Url,
		}).Info(infoData.Message)

		logrusLogger.WithFields(logrus.Fields{
			"Code": errorData.Code,
			"Err":  errorData.Err,
			"Url":  errorData.Url,
		}).Error(errorData.Message)
	}
}
