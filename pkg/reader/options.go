package reader

import (
	"fmt"

	"github.com/bom-squad/protobom/pkg/formats"
	"github.com/bom-squad/protobom/pkg/native"
)

type Options struct {
	Format             formats.Format
	UnserializeOptions *native.UnserializeOptions
	formatOptions      map[string]interface{}
}

// argToOptsKeyVal returns a key value to access the options dictionary by using
// key as a string or its type if its a serializer driver.
func argToOptsKeyVal(key interface{}) string {
	keyVal, ok := key.(string)
	if !ok {
		keyVal = fmt.Sprintf("%T", key)
	}

	return keyVal
}

func (o *Options) GetFormatOptions(key interface{}) interface{} {
	keyVal := argToOptsKeyVal(key)
	if _, ok := o.formatOptions[keyVal]; ok {
		return o.formatOptions[keyVal]
	}
	return nil
}

func (o *Options) SetFormatOptions(key, opts interface{}) {
	if o.formatOptions == nil {
		o.formatOptions = map[string]interface{}{}
	}
	keyVal := argToOptsKeyVal(key)
	if keyVal == "" {
		return
	}
	o.formatOptions[keyVal] = opts
}

type ReaderOption func(*Reader)

func WithFormatOptions(driverKey string, opts interface{}) ReaderOption {
	return func(r *Reader) {
		r.Options.SetFormatOptions(driverKey, opts)
	}
}

func WithUnserializeOptions(uo *native.UnserializeOptions) ReaderOption {
	return func(r *Reader) {
		if uo != nil {
			r.Options.UnserializeOptions = uo
		}
	}
}

func WithSniffer(s Sniffer) ReaderOption {
	return func(r *Reader) {
		if s != nil {
			r.sniffer = s
		}
	}
}
