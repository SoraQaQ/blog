package decorator

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
)

type commandLoggingDecorator[C any] struct {
	log  *log.Helper
	base CommandHandler[C]
}

func (d *commandLoggingDecorator[C]) Handler(ctx context.Context, cmd C) (err error) {
	// log.Infof("cmd: %v", cmd)
	body, _ := json.Marshal(cmd)
	d.log.WithContext(ctx).Infow("Command", "CommandName", generateActionName(cmd), "body", string(body))
	defer func() {
		if err == nil {
			d.log.WithContext(ctx).Infow("Command", "CommandName", generateActionName(cmd), "body", string(body))
		} else {
			d.log.WithContext(ctx).Errorf("command failed: %v", err)
		}
	}()
	err = d.base.Handler(ctx, cmd)
	return err
}

type queryLoggingDecorator[Q, R any] struct {
	log  *log.Helper
	base QueryHandler[Q, R]
}

func (d *queryLoggingDecorator[Q, R]) Handler(ctx context.Context, query Q) (result R, err error) {
	body, _ := json.Marshal(query)
	d.log.WithContext(ctx).Infow("Query", "QueryName", generateActionName(query), "body", string(body))
	defer func() {
		if err == nil {
			d.log.WithContext(ctx).Infow("Query", "QueryName", generateActionName(query), "body", string(body))
		} else {
			d.log.WithContext(ctx).Errorf("Query failed: %v", err)
		}
	}()
	result, err = d.base.Handler(ctx, query)
	return
}

func generateActionName(cmd any) string {
	return strings.Split(fmt.Sprintf("%T", cmd), ".")[1]
}
