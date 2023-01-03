package dkron

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/sohlich/elogrus.v7"
	"io/ioutil"
	"sync"
)

// ginOnce is a wrapper around gin global var changes. This is a workaround
// against the lack of concurrency safety of these vars in the gin package.
var ginOnce sync.Once

// InitLogger creates the logger instance
func InitLogger(logLevel string, node string) *logrus.Entry {
	formattedLogger := logrus.New()
	formattedLogger.Formatter = &logrus.TextFormatter{FullTimestamp: true}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithError(err).Error("Error parsing log level, using: info")
		level = logrus.InfoLevel
	}

	formattedLogger.Level = level
	log := logrus.NewEntry(formattedLogger).WithField("node", node)

	ginOnce.Do(func() {
		if level == logrus.DebugLevel {
			gin.DefaultWriter = log.Writer()
			gin.SetMode(gin.DebugMode)
		} else {
			gin.DefaultWriter = ioutil.Discard
			gin.SetMode(gin.ReleaseMode)
		}
	})

	vi := viper.New()
	vi.SetConfigName("dkron-es-config") // name of config file (without extension)
	vi.AddConfigPath("/etc/dkron")      // call multiple times to add many search paths
	vi.AddConfigPath("$HOME/.dkron")    // call multiple times to add many search paths
	vi.AddConfigPath("./config")        // call multiple times to add many search paths
	er := vi.ReadInConfig()             // Find and read the config file
	if er != nil {                      // Handle errors reading the config file
		log.WithError(er).Info("No valid elasticsearch config found")
	} else {
		esUrl := vi.GetString("es-url")
		if esUrl != "" {
			client, err := elastic.NewClient(elastic.SetURL(esUrl))
			if err != nil {
				log.Panic(err)
			}
			hook, err := elogrus.NewAsyncElasticHook(client, node, logrus.DebugLevel, "dkron-log")
			if err != nil {
				log.Panic(err)
			}
			formattedLogger.AddHook(hook)
		} else {
			log.WithError(er).Info("No valid elasticsearch config found")
		}
	}

	return log
}
