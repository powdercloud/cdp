// Code generated by cdpgen. DO NOT EDIT.

// Package domsnapshot implements the DOMSnapshot domain. This domain
// facilitates obtaining document snapshots with DOM, layout, and style
// information.
package domsnapshot

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the DOMSnapshot domain. This domain
// facilitates obtaining document snapshots with DOM, layout, and style
// information.
type domainClient struct {
	conn      *rpcc.Conn
	sessionID string
}

// NewClient returns a client for the DOMSnapshot domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// NewClient returns a client for the DOMSnapshot domain with the connection set to conn.
func NewSessionClient(conn *rpcc.Conn, sessionID string) *domainClient {
	return &domainClient{conn: conn, sessionID: sessionID}
}

// Disable invokes the DOMSnapshot method. Disables DOM snapshot agent for the
// given page.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.InvokeRPC(ctx, "DOMSnapshot.disable", d.sessionID, nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "DOMSnapshot", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the DOMSnapshot method. Enables DOM snapshot agent for the
// given page.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.InvokeRPC(ctx, "DOMSnapshot.enable", d.sessionID, nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "DOMSnapshot", Op: "Enable", Err: err}
	}
	return
}

// GetSnapshot invokes the DOMSnapshot method. Returns a document snapshot,
// including the full DOM tree of the root node (including iframes, template
// contents, and imported documents) in a flattened array, as well as layout
// and white-listed computed style information for the nodes. Shadow DOM in the
// returned DOM tree is flattened.
func (d *domainClient) GetSnapshot(ctx context.Context, args *GetSnapshotArgs) (reply *GetSnapshotReply, err error) {
	reply = new(GetSnapshotReply)
	if args != nil {
		err = rpcc.InvokeRPC(ctx, "DOMSnapshot.getSnapshot", d.sessionID, args, reply, d.conn)
	} else {
		err = rpcc.InvokeRPC(ctx, "DOMSnapshot.getSnapshot", d.sessionID, nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOMSnapshot", Op: "GetSnapshot", Err: err}
	}
	return
}

// CaptureSnapshot invokes the DOMSnapshot method. Returns a document
// snapshot, including the full DOM tree of the root node (including iframes,
// template contents, and imported documents) in a flattened array, as well as
// layout and white-listed computed style information for the nodes. Shadow DOM
// in the returned DOM tree is flattened.
func (d *domainClient) CaptureSnapshot(ctx context.Context, args *CaptureSnapshotArgs) (reply *CaptureSnapshotReply, err error) {
	reply = new(CaptureSnapshotReply)
	if args != nil {
		err = rpcc.InvokeRPC(ctx, "DOMSnapshot.captureSnapshot", d.sessionID, args, reply, d.conn)
	} else {
		err = rpcc.InvokeRPC(ctx, "DOMSnapshot.captureSnapshot", d.sessionID, nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "DOMSnapshot", Op: "CaptureSnapshot", Err: err}
	}
	return
}
