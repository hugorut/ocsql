package ocsql

// TraceOption allows for managing ocsql configuration using functional options.
type TraceOption func(o *TraceOptions)

// TraceOptions holds configuration of our ocsql tracing middleware.
// By default all options are set to false intentionally when creating a wrapped
// driver and provide the most sensible default with both performance and
// security in mind.
type TraceOptions struct {
	// AllowRoot, if set to true, will allow ocsql to create root spans in
	// absence of exisiting spans or even context.
	// Default is to not trace ocsql calls if no existing parent span is found
	// in context or when using methods not taking context.
	AllowRoot bool

	// Transaction, if set to true, will create spans for the duration of db
	// transactions. All spans created by the transaction's scoped queries will
	// become children of the transaction span.
	Transaction bool

	// Ping, if set to true, will enable the creation of spans on Ping requests.
	Ping bool

	// RowsNext, if set to true, will enable the creation of spans on RowsNext
	// calls. This can result in many spans.
	RowsNext bool

	// RowsClose, if set to true, will enable the creation of spans on RowsClose
	// calls.
	RowsClose bool

	// RowsAffected, if set to true, will enable the creation of spans on
	// RowsAffected calls.
	RowsAffected bool

	// LastInsertID, if set to true, will enable the creation of spans on
	// LastInsertId calls.
	LastInsertID bool

	// Query, if set to true, will enable recording of sql queries in spans.
	// Only allow this if it is safe to have queries recorded with respect to
	// security.
	Query bool

	// QueryParams, if set to true, will enable recording of parameters used
	// with parametrized queries. Only allow this if it is safe to have
	// parameters recorded with respect to security.
	// This setting is a noop if the Query option is set to false.
	QueryParams bool
}

// TraceAll has all tracing options enabled.
var TraceAll = TraceOptions{
	AllowRoot:    true,
	Transaction:  true,
	Ping:         true,
	RowsNext:     true,
	RowsClose:    true,
	RowsAffected: true,
	LastInsertID: true,
	Query:        true,
	QueryParams:  true,
}

// WithOptions sets our ocsql tracing middleware options through a single
// TraceOptions object.
func WithOptions(options TraceOptions) TraceOption {
	return func(o *TraceOptions) {
		*o = options
	}
}

// WithAllowRoot if set to true, will allow ocsql to create root spans in
// absence of exisiting spans or even context.
// Default is to not trace ocsql calls if no existing parent span is found
// in context or when using methods not taking context.
func WithAllowRoot(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.AllowRoot = b
	}
}

// WithTransaction if set to true, will create spans for the duration of db
// transactions. All spans created by the transaction's scoped queries will
// become children of the transaction span.
func WithTransaction(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Transaction = b
	}
}

// WithPing if set to true, will enable the creation of spans on Ping requests.
func WithPing(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Ping = b
	}
}

// WithRowsNext if set to true, will enable the creation of spans on RowsNext
// calls. This can result in many spans.
func WithRowsNext(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.RowsNext = b
	}
}

// WithRowsClose if set to true, will enable the creation of spans on RowsClose
// calls.
func WithRowsClose(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.RowsClose = b
	}
}

// WithRowsAffected if set to true, will enable the creation of spans on
// RowsAffected calls.
func WithRowsAffected(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.RowsAffected = b
	}
}

// WithLastInsertID if set to true, will enable the creation of spans on
// LastInsertId calls.
func WithLastInsertID(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.LastInsertID = b
	}
}

// WithQuery if set to true, will enable recording of sql queries in spans.
// Only allow this if it is safe to have queries recorded with respect to
// security.
func WithQuery(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Query = b
	}
}

// WithQueryParams if set to true, will enable recording of parameters used
// with parametrized queries. Only allow this if it is safe to have
// parameters recorded with respect to security.
// This setting is a noop if the Query option is set to false.
func WithQueryParams(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.QueryParams = b
	}
}
