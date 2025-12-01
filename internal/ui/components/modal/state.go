package modal

type OverlayType int

const (
	OverlayNone OverlayType = iota
	OverlaySearch
	OverlayEditList
	OverlayEditInput
	OverlayEditBool
)
