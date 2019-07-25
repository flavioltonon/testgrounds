package testgrounds

const OPTION_TYPE_BUCKET = "bucket-option"

func init() {
	RegisterOption(&BucketOption{})
}

type BucketOption struct {
	Name    string `json:"name"`
	Project string `json:"project"`
}

func (o BucketOption) Type() string {
	return OPTION_TYPE_BUCKET
}

type BucketOptions []BucketOption
