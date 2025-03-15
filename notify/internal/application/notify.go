/*
Bot Service
*/
package application

import (
	"context"

	"golang.org/x/sync/errgroup"

	"github.com/shortlink-org/shortlink/boundaries/notification/notify/internal/domain/events"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/notify"
)

func New(dataBus mq.MQ, log logger.Logger) (*Bot, error) {
	return &Bot{
		mq:  dataBus,
		log: log,
	}, nil
}

func (b *Bot) Use(ctx context.Context) {
	// Subscribe to Event
	notify.Subscribe(events.METHOD_NEW_LINK, b)

	// TODO: refactoring this code
	// getEventNewLink := query.Response{
	// 	Chan: make(chan query.ResponseMessage),
	// }

	g := errgroup.Group{}

	// Subscribe to MQ Event
	// g.Go(func() error {
	// 	if b.mq != nil {
	// 		if errSubscribe := b.mq.Subscribe(ctx, link.MQ_EVENT_LINK_CREATED, getEventNewLink); errSubscribe != nil {
	// 			return errSubscribe
	// 		}
	// 	}
	//
	// 	return nil
	// })

	// Listen to MQ Event
	// g.Go(func() error {
	// 	for {
	// 		msg := <-getEventNewLink.Chan
	//
	// 		// Convert: []byte to link.Link
	// 		myLink := &link.Link{}
	// 		if err := proto.Unmarshal(msg.Body, myLink); err != nil {
	// 			b.log.ErrorWithContext(msg.Context, fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
	// 			continue
	// 		}
	//
	// 		b.log.InfoWithContext(msg.Context, "Get new LINK", field.Fields{"url": myLink.GetUrl()})
	// 		notify.Publish(msg.Context, events.METHOD_NEW_LINK, myLink, nil)
	// 	}
	// })

	// Wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		b.log.Error(err.Error())
	}
}

// Notify - Notify to Bot
func (b *Bot) Notify(ctx context.Context, event uint32, payload any) notify.Response[any] {
	// switch event {
	// case events.METHOD_NEW_LINK:
	// 	if addLink, ok := payload.(*link.Link); ok {
	// 		b.send(ctx, addLink)
	// 	}
	// }

	return notify.Response[any]{}
}

// func (b *Bot) send(ctx context.Context, in *link.Link) {
func (b *Bot) send(ctx context.Context, in any) {
	// payload := fmt.Sprintf("LINK: %s", in.GetUrl())
	//
	// notify.Publish(ctx, events.METHOD_SEND_NEW_LINK, payload, nil)
}
