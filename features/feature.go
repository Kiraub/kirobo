package features

// FeatureKey is used to enable and disable features
type FeatureKey uint8

// Feature ...
type Feature interface {
	IsEnabled() bool
	Enable(func())
	Disable()
	Handler() interface{}
}

type feature struct {
	enabled  bool
	handler  interface{}
	disabler func()
}

// CreateFeature packs a handler function into an internal struct that implements the Feature interface
func CreateFeature(handler interface{}) Feature {
	f := new(feature)
	f.enabled = false
	f.handler = handler
	f.disabler = nil
	return f
}

// IsEnabled returns whether the feature is enabled
func (r *feature) IsEnabled() bool {
	return r.enabled
}

// Enable ...
func (r *feature) Enable(disabler func()) {
	if r.IsEnabled() {
		r.Disable()
	}
	r.enabled = true
	r.disabler = disabler
}

// Disable ...
func (r *feature) Disable() {
	if !r.IsEnabled() {
		return
	}
	r.disabler()
	r.enabled = false
	r.disabler = nil
}

// Handler ...
func (r *feature) Handler() interface{} {
	return r.handler
}
