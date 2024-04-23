package logger

import (
	"io"
	"os"
	"personal_page/mq"

	"github.com/spf13/viper"
)

type LoggerOut struct {
	mq *mq.RabbitMQ
}

func NewLoggerOut(conf *viper.Viper, mq *mq.RabbitMQ) *LoggerOut {
	out := &LoggerOut{}
	if conf.GetBool("log.mq") {
		out.mq = mq
	}
	return out
}

// Write implements io.Writer.
func (o *LoggerOut) Write(p []byte) (n int, err error) {
	if o.mq != nil {
		err := o.mq.Send(p)
		if err != nil {
			return 0, err
		}
	}
	return os.Stdout.Write(p)
}

var _ io.Writer = &LoggerOut{}
