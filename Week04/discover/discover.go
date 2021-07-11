package discover

import (
	"context"
	"fmt"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/discovery/consul"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"time"
)

func Discover() {
	conf := &consul.SDConfig{Services: []string{"configuredServiceName"}, RefreshInterval: model.Duration(1 * time.Millisecond)}
	d, err := consul.NewDiscovery(conf, nil)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	ch := make(chan []*targetgroup.Group)

	go d.Run(context.Background(), ch)

	fmt.Printf("%+v", <-ch)
}
