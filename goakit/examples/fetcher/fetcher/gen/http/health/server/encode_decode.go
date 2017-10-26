// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// health HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/client/fetcher/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/http"
)

// EncodeShowResponse returns an encoder for responses returned by the health
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}