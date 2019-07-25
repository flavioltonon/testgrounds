package testgrounds

type Client interface {
	NewBucket(options ...Option) (Bucket, error)
}
