package urlquery

type options struct {
	urlEncoder       UrlEncoder
	needEmptyValue bool
}

type Option interface {
	apply(*options)
}

type urlEncoderOption struct {
	urlEncoder UrlEncoder
}

func (o urlEncoderOption) apply(opts *options) {
	opts.urlEncoder = o.urlEncoder
}

//support customized urlEncoder option
func WithUrlEncoder(u UrlEncoder) Option {
	return urlEncoderOption{urlEncoder: u}
}

type NeedEmptyValueOption bool

func (o NeedEmptyValueOption) apply(opts *options) {
	opts.needEmptyValue = bool(o)
}

//support to control whether to ignore empty value.
//It just happen to the element directly in strcut, not include map slice array
//default:false, meaning not to ignore
func WithNeedEmptyValue(c bool) Option {
	return NeedEmptyValueOption(c)
}
