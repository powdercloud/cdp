// Code generated by cdpgen. DO NOT EDIT.

// Package console implements the Console domain. This domain is deprecated -
// use Runtime or Log instead.
package console

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Console domain. This domain is deprecated
// - use Runtime or Log instead.
type domainClient struct {
	conn      *rpcc.Conn
	sessionID string
}

// NewClient returns a client for the Console domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// NewClient returns a client for the Console domain with the connection set to conn.
func NewSessionClient(conn *rpcc.Conn, sessionID string) *domainClient {
	return &domainClient{conn: conn, sessionID: sessionID}
}

// ClearMessages invokes the Console method. Does nothing.
func (d *domainClient) ClearMessages(ctx context.Context) (err error) {
	err = rpcc.InvokeRPC(ctx, "Console.clearMessages", d.sessionID, nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Console", Op: "ClearMessages", Err: err}
	}
	return
}

// Disable invokes the Console method. Disables console domain, prevents
// further console messages from being reported to the client.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.InvokeRPC(ctx, "Console.disable", d.sessionID, nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Console", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the Console method. Enables console domain, sends the
// messages collected so far to the client by means of the `messageAdded`
// notification.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.InvokeRPC(ctx, "Console.enable", d.sessionID, nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Console", Op: "Enable", Err: err}
	}
	return
}

func (d *domainClient) MessageAdded(ctx context.Context) (MessageAddedClient, error) {
	s, err := rpcc.NewStream(ctx, "Console.messageAdded", d.sessionID, d.conn)
	if err != nil {
		return nil, err
	}
	return &messageAddedClient{Stream: s}, nil
}

type messageAddedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *messageAddedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *messageAddedClient) Recv() (*MessageAddedReply, error) {
	event := new(MessageAddedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Console", Op: "MessageAdded Recv", Err: err}
	}
	return event, nil
}
