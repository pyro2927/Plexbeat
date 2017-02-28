package beater

import (
	"fmt"
	"time"
  "strings"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

  "github.com/jrudio/go-plex-client"

	"github.com/pyro2927/plexbeat/config"
)

type Plexbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Plexbeat{
		done: make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Plexbeat) Run(b *beat.Beat) error {
	logp.Info("plexbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
  s := []string{"http://", bt.config.Host, ":", bt.config.Port};
  Plex, _ := plex.New(strings.Join(s, ""), bt.config.AuthToken)
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}
    current, err := Plex.GetSessions()

    if err == nil {
      event := common.MapStr{
        "@timestamp":             common.Time(time.Now()),
        "type":                   b.Name,
        "counter":                counter,
        "plex.sessions.count":    current.Size,
        "plex.host":              bt.config.Host,
      }
      bt.client.PublishEvent(event)
      logp.Info("Event sent")
      counter++
    } else {
      logp.Warn("Unable to connect to Plex Server")
    }
	}
}

func (bt *Plexbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
