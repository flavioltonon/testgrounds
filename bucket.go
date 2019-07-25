package testgrounds

type Bucket interface {
	Create() error
	Attributes() (BucketAttributes, error)
}
