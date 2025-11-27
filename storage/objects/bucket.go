package objects

import (
	"context"
	"errors"
	"iter"
	"net/url"
	"os"
)

// Bucket represents an object storage bucket, containing a set of files.
//
// See NewBucket for more information on how to declare a Bucket.
type Bucket struct {
	_ int // for godoc to show unexported fields
}

// BucketConfig is the configuration for a Bucket.
type BucketConfig struct {
	// Whether objects in the bucket should be publicly accessible, via CDN.
	Public bool

	// Whether objects stored in the bucket should be versioned.
	//
	// If true, the bucket will store multiple versions of each object
	// whenever it changes, as opposed to overwriting the old version.
	Versioned bool
}

// Upload uploads a new object to the bucket.
//
// The returned writer must be successfully closed for the upload to complete.
// To abort the upload, call (*Writer).Abort or cancel the provided context.
func (*Bucket) Upload(ctx context.Context, object string, options ...UploadOption) (_ *Writer) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L99-L131
	doPanic("encore apps must be run using the encore command")
	return
}

// PublicURL returns the public URL for accessing an object in the bucket.
func (*Bucket) PublicURL(object string, options ...PublicURLOption) (_ *url.URL) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L134-L150
	doPanic("encore apps must be run using the encore command")
	return
}

// Writer is the writer for an object being uploaded to a bucket.
type Writer struct {
	_ int // for godoc to show unexported fields
}

// Write writes data to the object being uploaded.
func (*Writer) Write(p []byte) (_ int, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L170-L173
	doPanic("encore apps must be run using the encore command")
	return
}

// Abort aborts the upload.
func (*Writer) Abort(err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L176-L182
	doPanic("encore apps must be run using the encore command")
	return
}

// Close closes the upload, completing the upload if no errors occurred.
func (*Writer) Close() (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L185-L208
	doPanic("encore apps must be run using the encore command")
	return
}

// Download downloads an object from the bucket.
// Any error is encountered is reported by the methods on *Reader.
// To check if the operation failed, call (*Reader).Err.
//
// If the object does not exist, the error may be checked with errors.Is(err, ErrObjectNotFound).
func (*Bucket) Download(ctx context.Context, object string, options ...DownloadOption) (_ *Reader) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L249-L277
	doPanic("encore apps must be run using the encore command")
	return
}

// Reader is the reader for an object being downloaded from a bucket.
type Reader struct {
	_ int // for godoc to show unexported fields
}

// Err returns the error encountered during reading, if any.
func (*Reader) Err() (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L292-L294
	doPanic("encore apps must be run using the encore command")
	return
}

// Read reads data from the object being downloaded.
func (*Reader) Read(p []byte) (_ int, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L297-L306
	doPanic("encore apps must be run using the encore command")
	return
}

// Close closes the reader.
// It must be called to release resources.
func (*Reader) Close() (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L310-L318
	doPanic("encore apps must be run using the encore command")
	return
}

// Query describes the set of objects to query for using List.
type Query struct {
	// Prefix indicates to only return objects
	// whose name starts with the given prefix.
	Prefix string

	// Maximum number of objects to return. Zero means no limit.
	Limit int64
}

// ObjectAttrs describes the attributes of an object.
type ObjectAttrs struct {
	// The name of the object.
	Name string

	// The version of the object, if bucket versioning is enabled.
	Version string

	// The content type of the object, if set during upload.
	ContentType string

	// The size of the object, in bytes.
	Size int64

	// The computed ETag of the object.
	ETag string
}

// ListEntry describes an objects during listing.
type ListEntry struct {
	// The name of the object.
	Name string
	// The size of the object, in bytes.
	Size int64
	// The computed ETag of the object.
	ETag string
}

type SignedUploadURL struct {
	// The signed URL
	URL string
}

type SignedDownloadURL struct {
	// The signed URL
	URL string
}

// List lists objects in the bucket.
func (*Bucket) List(ctx context.Context, query *Query, options ...ListOption) (_ iter.Seq2[*ListEntry, error]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L415-L467
	doPanic("encore apps must be run using the encore command")
	return
}

// Remove removes an object from the bucket.
func (*Bucket) Remove(ctx context.Context, object string, options ...RemoveOption) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L470-L513
	doPanic("encore apps must be run using the encore command")
	return
}

var (
	// ErrObjectNotFound is returned when requested object does not exist in the bucket.
	ErrObjectNotFound = errors.New("objects: object doesn't exist")

	// ErrPreconditionFailed is returned when a precondition for an operation is not met,
	// such as when an object already exists and Preconditions.NotExists is true.
	ErrPreconditionFailed = errors.New("objects: precondition failed")

	// ErrInvalidArgument is returned when an argument for an operation is invalid or out
	// of bounds. Such as when a too long time-to-live is passed to a sign URL operation.
	ErrInvalidArgument = errors.New("objects: invalid argument")
)

// Attrs returns the attributes of an object in the bucket.
// If the object does not exist, it returns ErrObjectNotFound.
func (*Bucket) Attrs(ctx context.Context, object string, options ...AttrsOption) (_ *ObjectAttrs, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L530-L588
	doPanic("encore apps must be run using the encore command")
	return
}

// Generates an external URL to allow uploading an object to the bucket.
//
// Anyone with possession of the URL can write to the given object name
// without any additional auth.
func (*Bucket) SignedUploadURL(ctx context.Context, object string, options ...UploadURLOption) (_ *SignedUploadURL, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L594-L614
	doPanic("encore apps must be run using the encore command")
	return
}

// Generates an external URL to allow downloading an object from the bucket.
//
// Anyone with possession of the URL can read the given object
// without any additional auth.
func (*Bucket) SignedDownloadURL(ctx context.Context, object string, options ...DownloadURLOption) (_ *SignedDownloadURL, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L620-L640
	doPanic("encore apps must be run using the encore command")
	return
}

// Exists reports whether an object exists in the bucket.
func (*Bucket) Exists(ctx context.Context, object string, options ...ExistsOption) (_ bool, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/objects/bucket.go#L643-L702
	doPanic("encore apps must be run using the encore command")
	return
}

// doPanic is a wrapper around panic to prevent static analysis tools
// from thinking Encore APIs unconditionally panic.,
func doPanic(v any) {
	if os.Getenv("ENCORERUNTIME_NOPANIC") == "" {
		panic(v)
	}
}
