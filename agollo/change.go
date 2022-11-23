package agollo

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4/storage"
)

type CustomChangeListener struct {
}

func (c *CustomChangeListener) OnChange(changeEvent *storage.ChangeEvent) {
	//打印被更改的key
	for key, value := range changeEvent.Changes {
		fmt.Println(changeEvent.Namespace+" - change key : ", key, ", value :", value)
	}
}

func (c *CustomChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	//打印全部key
}
