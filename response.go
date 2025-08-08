package spsp

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-http500"
)

// Response represents a Simple Payment Setup Protocol (SPSP) reponse, as defined here:
// https://interledger.org/developers/rfcs/simple-payment-setup-protocol/
//
// The DestinationAccount (destination_account) is an ILP Address of the SPSP Server.
// In case of push payments, this is the receiver.
// In case of pull payments, this is the sender.
//
// The SharedSecret (shared_secret) is the base64 encoding of 32 bytes.
// The shared secret to be used by this specific HTTP client in the STREAM.
// Should be shared only by the server and this specific HTTP client, and should therefore be different in each query response.
// Even though clients SHOULD accept base64url encoded secrets, base64 encoded secrets are recommended.
// Use [Response.EncodeAndSetSharedSecret] to set SharedSecret.
// Use [Response.DecodeSharedSecret] to get the SharedSecret.
//
// The ReceiptsEnabled (receipts_enabled) field is optional.
// If true, the SPSP server will issue STREAM Receipts in the STREAM connection.
// If false or omitted, the server will not issue STREAM Receipts.
type Response struct {
	DestinationAccount string `json:"destination_account"`
	SharedSecret       string `json:"shared_secret"`
	ReceiptsEnabled    bool   `json:"receipts_enabled"`
}

// EncodeAndSetSharedSecret base64 encods the 32 byte shared-secret and sets it
// as the value of the value of the Response.SharedSecret field.
func (receiver *Response) EncodeAndSetSharedSecret(src [SharedSecretDecodedLength]byte) {
	if nil == receiver {
		return
	}

	receiver.SharedSecret = base64.StdEncoding.EncodeToString(src[:])
}

// DecodeSharedSecret base64 decodes the value of the Response.SharedSecret field into a [32]byte.
func (receiver Response) DecodeSharedSecret(dst *[SharedSecretDecodedLength]byte) error {
	if nil == dst {
		return errNilDestination
	}

	decoded, err := base64.StdEncoding.DecodeString(receiver.SharedSecret)
	if nil != err {
		return err
	}

	if SharedSecretDecodedLength != len(decoded) {
		return erorr.Errorf("spsp: ")
	}

	copy(dst[:], decoded)
	return nil
}

// ServeHTTP makes it so [Response] is also an [http.Handler], and thus — it can send itself
// as an HTTP response.
//
// Note that ServeHTTP also sets the appropriate HTTP response headers specified by:
// https://interledger.org/developers/rfcs/simple-payment-setup-protocol/
func (receiver Response) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if nil == responseWriter {
		return
	}
	if nil == request {
		http500.InternalServerError(responseWriter, request)
		return
	}

	responseWriter.Header().Set("Access-Control-Allow-Headers", "web-monetization-id")
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Cache-Control", "no-cache")
	responseWriter.Header().Set("Content-Type", "application/spsp4+json")

	var builder strings.Builder

	err := json.NewEncoder(responseWriter).Encode(&builder)
	if nil != err {
		http500.InternalServerError(responseWriter, request)
		return
	}

	io.WriteString(responseWriter, builder.String())
}
