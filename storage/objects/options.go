package objects

import (
	"time"
)

// DownloadOption describes available options for the Download operation.
type DownloadOption interface {
	downloadOption()
}

// WithVersion specifies that the operation should be performed
// against the provided version of the object.
func WithVersion(version string) (_ withVersionOption) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/storage/objects/options.go#L19-L21
	doPanic("encore apps must be run using the encore command")
	return
}

type withVersionOption struct {
	version string
}

func (o withVersionOption) downloadOption() {}

func (o withVersionOption) removeOption() {}

func (o withVersionOption) attrsOption() {}

func (o withVersionOption) existsOption() {}

func (o withTTLOption) uploadURLOption() {}

func (o withTTLOption) downloadURLOption() {}

// WithTTL is used for signed URLs, to specify the lifetime of the generated
// URL. The max value is seven days. The default lifetime, if this
// option is missing, is one hour.
func WithTTL(TTL time.Duration) (_ withTTLOption) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/storage/objects/options.go#L56-L58
	doPanic("encore apps must be run using the encore command")
	return
}

type withTTLOption struct {
	TTL time.Duration
}

type downloadOptions struct {
	version string
}

// UploadOption describes available options for the Upload operation.
type UploadOption interface {
}

// WithPreconditions is an UploadOption for only uploading an object
// if certain preconditions are met.
func WithPreconditions(pre Preconditions) (_ withPreconditionsOption) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/storage/objects/options.go#L79-L81
	doPanic("encore apps must be run using the encore command")
	return
}

// Preconditions are the available preconditions for an upload operation.
type Preconditions struct {
	// NotExists specifies that the object must not exist prior to uploading.
	NotExists bool
}

type withPreconditionsOption struct {
	pre Preconditions
}

func (o withPreconditionsOption) uploadOption() {}

// UploadAttrs specifies additional object attributes to set during upload.
type UploadAttrs struct {
	// ContentType specifies the content type of the object.
	ContentType string
}

// WithUploadAttrs is an UploadOption for specifying additional object attributes
// to set during upload.
func WithUploadAttrs(attrs UploadAttrs) (_ withUploadAttrsOption) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/storage/objects/options.go#L109-L111
	doPanic("encore apps must be run using the encore command")
	return
}

type withUploadAttrsOption struct {
	attrs UploadAttrs
}

func (o withUploadAttrsOption) uploadOption() {}

// ListOption describes available options for the List operation.
type ListOption interface {
	listOption()
}

// RemoveOption describes available options for the Remove operation.
type RemoveOption interface {
	removeOption()
}

// AttrsOption describes available options for the Attrs operation.
type AttrsOption interface {
	attrsOption()
}

// UploadURLOption describes available options for the SignedUploadURL operation.
type UploadURLOption interface {
	uploadURLOption()
}

// DownloadURLOption describes available options for the SignedDownloadURL operation.
type DownloadURLOption interface {
	downloadURLOption()
}

// ExistsOption describes available options for the Exists operation.
type ExistsOption interface {
	existsOption()
}

// PublicURLOption describes available options for the PublicURL operation.
type PublicURLOption interface {
	publicURLOption()
}
