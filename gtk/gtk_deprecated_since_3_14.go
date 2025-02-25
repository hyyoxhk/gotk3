// +build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_deprecated

package gtk

// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_deprecated_since_3_14.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_alignment_get_type()), marshalAlignment},
		{glib.Type(C.gtk_arrow_get_type()), marshalArrow},
		{glib.Type(C.gtk_misc_get_type()), marshalMisc},
		{glib.Type(C.gtk_status_icon_get_type()), marshalStatusIcon},
	}
	glib.RegisterGValueMarshalers(tm)

	//Contribute to casting
	for k, v := range map[string]WrapFn{
		"GtkAlignment":  wrapAlignment,
		"GtkArrow":      wrapArrow,
		"GtkMisc":       wrapMisc,
		"GtkStatusIcon": wrapStatusIcon,
	} {
		WrapMap[k] = v
	}
}

/*
 * deprecated since version 3.14 and should not be used in newly-written code
 */

/*
 * GtkTreeView
 */

// TODO:
// gtk_tree_view_set_rules_hint().
// gtk_tree_view_get_rules_hint().

/*
 * GtkWindow
 */

// SetHasResizeGrip is a wrapper around gtk_window_set_has_resize_grip().
func (v *Window) SetHasResizeGrip(setting bool) {
	C.gtk_window_set_has_resize_grip(v.native(), gbool(setting))
}

// GetHasResizeGrip is a wrapper around gtk_window_get_has_resize_grip().
func (v *Window) GetHasResizeGrip() bool {
	c := C.gtk_window_get_has_resize_grip(v.native())
	return gobool(c)
}

// ResizeGripIsVisible is a wrapper around gtk_window_resize_grip_is_visible().
func (v *Window) ResizeGripIsVisible() bool {
	c := C.gtk_window_resize_grip_is_visible(v.native())
	return gobool(c)
}

// GetResizeGripArea is a wrapper around gtk_window_get_resize_grip_area().
func (v *Window) GetResizeGripArea() (*gdk.Rectangle, bool) {
	var cRect *C.GdkRectangle
	wasRetrieved := C.gtk_window_get_resize_grip_area(v.native(), cRect)
	rect := gdk.WrapRectangle(uintptr(unsafe.Pointer(cRect)))
	return rect, gobool(wasRetrieved)
}

/*
 * GtkWidget
 */

// Reparent() is a wrapper around gtk_widget_reparent().
func (v *Widget) Reparent(newParent IWidget) {
	C.gtk_widget_reparent(v.native(), newParent.toWidget())
}

// SetDoubleBuffered is a wrapper around gtk_widget_set_double_buffered().
func (v *Widget) SetDoubleBuffered(doubleBuffered bool) {
	C.gtk_widget_set_double_buffered(v.native(), gbool(doubleBuffered))
}

// GetDoubleBuffered is a wrapper around gtk_widget_get_double_buffered().
func (v *Widget) GetDoubleBuffered() bool {
	c := C.gtk_widget_get_double_buffered(v.native())
	return gobool(c)
}

// TODO:
// gtk_widget_region_intersect().

// GetPadding is a wrapper around gtk_alignment_get_padding().
func (v *Alignment) GetPadding() (top, bottom, left, right uint) {
	var ctop, cbottom, cleft, cright C.guint
	C.gtk_alignment_get_padding(v.native(), &ctop, &cbottom, &cleft,
		&cright)
	return uint(ctop), uint(cbottom), uint(cleft), uint(cright)
}

// SetPadding is a wrapper around gtk_alignment_set_padding().
func (v *Alignment) SetPadding(top, bottom, left, right uint) {
	C.gtk_alignment_set_padding(v.native(), C.guint(top), C.guint(bottom),
		C.guint(left), C.guint(right))
}

// AlignmentNew is a wrapper around gtk_alignment_new().
func AlignmentNew(xalign, yalign, xscale, yscale float32) (*Alignment, error) {
	c := C.gtk_alignment_new(C.gfloat(xalign), C.gfloat(yalign), C.gfloat(xscale),
		C.gfloat(yscale))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAlignment(obj), nil
}

// Set is a wrapper around gtk_alignment_set().
func (v *Alignment) Set(xalign, yalign, xscale, yscale float32) {
	C.gtk_alignment_set(v.native(), C.gfloat(xalign), C.gfloat(yalign),
		C.gfloat(xscale), C.gfloat(yscale))
}

/*
 * GtkArrow
 */

// Arrow is a representation of GTK's GtkArrow.
type Arrow struct {
	Misc
}

// ArrowNew is a wrapper around gtk_arrow_new().
func ArrowNew(arrowType ArrowType, shadowType ShadowType) (*Arrow, error) {
	c := C.gtk_arrow_new(C.GtkArrowType(arrowType),
		C.GtkShadowType(shadowType))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapArrow(obj), nil
}

// Set is a wrapper around gtk_arrow_set().
func (v *Arrow) Set(arrowType ArrowType, shadowType ShadowType) {
	C.gtk_arrow_set(v.native(), C.GtkArrowType(arrowType), C.GtkShadowType(shadowType))
}

// SetAlignment() is a wrapper around gtk_button_set_alignment().
func (v *Button) SetAlignment(xalign, yalign float32) {
	C.gtk_button_set_alignment(v.native(), (C.gfloat)(xalign),
		(C.gfloat)(yalign))
}

// GetAlignment() is a wrapper around gtk_button_get_alignment().
func (v *Button) GetAlignment() (xalign, yalign float32) {
	var x, y C.gfloat
	C.gtk_button_get_alignment(v.native(), &x, &y)
	return float32(x), float32(y)
}

// SetReallocateRedraws is a wrapper around gtk_container_set_reallocate_redraws().
func (v *Container) SetReallocateRedraws(needsRedraws bool) {
	C.gtk_container_set_reallocate_redraws(v.native(), gbool(needsRedraws))
}

// GetAlignment is a wrapper around gtk_misc_get_alignment().
func (v *Misc) GetAlignment() (xAlign, yAlign float32) {
	var x, y C.gfloat
	C.gtk_misc_get_alignment(v.native(), &x, &y)
	return float32(x), float32(y)
}

// SetAlignment is a wrapper around gtk_misc_set_alignment().
func (v *Misc) SetAlignment(xAlign, yAlign float32) {
	C.gtk_misc_set_alignment(v.native(), C.gfloat(xAlign), C.gfloat(yAlign))
}

// GetPadding is a wrapper around gtk_misc_get_padding().
func (v *Misc) GetPadding() (xpad, ypad int) {
	var x, y C.gint
	C.gtk_misc_get_padding(v.native(), &x, &y)
	return int(x), int(y)
}

// SetPadding is a wrapper around gtk_misc_set_padding().
func (v *Misc) SetPadding(xPad, yPad int) {
	C.gtk_misc_set_padding(v.native(), C.gint(xPad), C.gint(yPad))
}

/*
 * GtkArrow
 * deprecated since version 3.14
 */
// native returns a pointer to the underlying GtkButton.
func (v *Arrow) native() *C.GtkArrow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkArrow(p)
}

func marshalArrow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapArrow(obj), nil
}

func wrapArrow(obj *glib.Object) *Arrow {
	if obj == nil {
		return nil
	}

	return &Arrow{Misc{Widget{glib.InitiallyUnowned{obj}}}}
}

/*
 * GtkAlignment
 * deprecated since version 3.14
 */

type Alignment struct {
	Bin
}

// native returns a pointer to the underlying GtkAlignment.
func (v *Alignment) native() *C.GtkAlignment {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkAlignment(p)
}

func marshalAlignment(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAlignment(obj), nil
}

func wrapAlignment(obj *glib.Object) *Alignment {
	if obj == nil {
		return nil
	}

	return &Alignment{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

/*
 * GtkStatusIcon
 * deprecated since version 3.14
 */

// StatusIcon is a representation of GTK's GtkStatusIcon.
// Deprecated since 3.14 in favor of notifications
// (no replacement, see https://stackoverflow.com/questions/41917903/gtk-3-statusicon-replacement)
type StatusIcon struct {
	*glib.Object
}

func marshalStatusIcon(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

func wrapStatusIcon(obj *glib.Object) *StatusIcon {
	if obj == nil {
		return nil
	}

	return &StatusIcon{obj}
}

func (v *StatusIcon) native() *C.GtkStatusIcon {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkStatusIcon(p)
}

// TODO: GtkStatusIcon * gtk_status_icon_new_from_gicon (GIcon *icon);
// TODO: void gtk_status_icon_set_from_gicon (GtkStatusIcon *status_icon, GIcon *icon);

// TODO: GIcon * gtk_status_icon_get_gicon (GtkStatusIcon *status_icon);

// TODO: void gtk_status_icon_set_screen (GtkStatusIcon *status_icon, GdkScreen *screen);
// TODO: GdkScreen * gtk_status_icon_get_screen (GtkStatusIcon *status_icon);

// TODO: GdkPixbuf * gtk_status_icon_get_pixbuf (GtkStatusIcon *status_icon);

// TODO: void gtk_status_icon_position_menu (GtkMenu *menu, gint *x, gint *y, gboolean *push_in, gpointer user_data);
// TODO: gboolean gtk_status_icon_get_geometry (GtkStatusIcon *status_icon, GdkScreen **screen, GdkRectangle *area, GtkOrientation *orientation);

// StatusIconNew is a wrapper around gtk_status_icon_new()
func StatusIconNew() (*StatusIcon, error) {
	c := C.gtk_status_icon_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapStatusIcon(glib.Take(unsafe.Pointer(c))), nil
}

// StatusIconNewFromFile is a wrapper around gtk_status_icon_new_from_file()
func StatusIconNewFromFile(filename string) (*StatusIcon, error) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_status_icon_new_from_file((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapStatusIcon(glib.Take(unsafe.Pointer(c))), nil
}

// StatusIconNewFromIconName is a wrapper around gtk_status_icon_new_from_icon_name()
func StatusIconNewFromIconName(iconName string) (*StatusIcon, error) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_status_icon_new_from_icon_name((*C.gchar)(cstr))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapStatusIcon(glib.Take(unsafe.Pointer(c))), nil
}

// StatusIconNewFromPixbuf is a wrapper around gtk_status_icon_new_from_pixbuf().
func StatusIconNewFromPixbuf(pixbuf *gdk.Pixbuf) (*StatusIcon, error) {
	c := C.gtk_status_icon_new_from_pixbuf(C.toGdkPixbuf(unsafe.Pointer(pixbuf.Native())))
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapStatusIcon(obj), nil
}

// SetFromFile is a wrapper around gtk_status_icon_set_from_file()
func (v *StatusIcon) SetFromFile(filename string) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_file(v.native(), (*C.gchar)(cstr))
}

// SetFromIconName is a wrapper around gtk_status_icon_set_from_icon_name()
func (v *StatusIcon) SetFromIconName(iconName string) {
	cstr := C.CString(iconName)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_from_icon_name(v.native(), (*C.gchar)(cstr))
}

// SetFromPixbuf is a wrapper around gtk_status_icon_set_from_pixbuf()
func (v *StatusIcon) SetFromPixbuf(pixbuf *gdk.Pixbuf) {
	C.gtk_status_icon_set_from_pixbuf(v.native(), C.toGdkPixbuf(unsafe.Pointer(pixbuf.Native())))
}

// GetStorageType is a wrapper around gtk_status_icon_get_storage_type()
func (v *StatusIcon) GetStorageType() ImageType {
	return (ImageType)(C.gtk_status_icon_get_storage_type(v.native()))
}

// SetTooltipText is a wrapper around gtk_status_icon_set_tooltip_text()
func (v *StatusIcon) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// GetTooltipText is a wrapper around gtk_status_icon_get_tooltip_text()
func (v *StatusIcon) GetTooltipText() string {
	c := C.gtk_status_icon_get_tooltip_text(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// SetTooltipMarkup is a wrapper around gtk_status_icon_set_tooltip_markup()
func (v *StatusIcon) SetTooltipMarkup(markup string) {
	cstr := C.CString(markup)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_tooltip_markup(v.native(), (*C.gchar)(cstr))
}

// GetTooltipMarkup is a wrapper around gtk_status_icon_get_tooltip_markup()
func (v *StatusIcon) GetTooltipMarkup() string {
	c := C.gtk_status_icon_get_tooltip_markup(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// SetHasTooltip is a wrapper around gtk_status_icon_set_has_tooltip()
func (v *StatusIcon) SetHasTooltip(hasTooltip bool) {
	C.gtk_status_icon_set_has_tooltip(v.native(), gbool(hasTooltip))
}

// GetTitle is a wrapper around gtk_status_icon_get_title()
func (v *StatusIcon) GetTitle() string {
	c := C.gtk_status_icon_get_title(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// SetName is a wrapper around gtk_status_icon_set_name()
func (v *StatusIcon) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_name(v.native(), (*C.gchar)(cstr))
}

// SetVisible is a wrapper around gtk_status_icon_set_visible()
func (v *StatusIcon) SetVisible(visible bool) {
	C.gtk_status_icon_set_visible(v.native(), gbool(visible))
}

// GetVisible is a wrapper around gtk_status_icon_get_visible()
func (v *StatusIcon) GetVisible() bool {
	return gobool(C.gtk_status_icon_get_visible(v.native()))
}

// IsEmbedded is a wrapper around gtk_status_icon_is_embedded()
func (v *StatusIcon) IsEmbedded() bool {
	return gobool(C.gtk_status_icon_is_embedded(v.native()))
}

// GetHasTooltip is a wrapper around gtk_status_icon_get_has_tooltip()
func (v *StatusIcon) GetHasTooltip() bool {
	return gobool(C.gtk_status_icon_get_has_tooltip(v.native()))
}

// SetTitle is a wrapper around gtk_status_icon_set_title()
func (v *StatusIcon) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_status_icon_set_title(v.native(), (*C.gchar)(cstr))
}

// GetIconName is a wrapper around gtk_status_icon_get_icon_name()
func (v *StatusIcon) GetIconName() string {
	c := C.gtk_status_icon_get_icon_name(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// GetSize is a wrapper around gtk_status_icon_get_size()
func (v *StatusIcon) GetSize() int {
	return int(C.gtk_status_icon_get_size(v.native()))
}

// PopupAtStatusIcon() is a wrapper around gtk_menu_popup() specific to usage with GtkStatusIcon.
// gomenu_popup() is defined in menu.go.h, this is a workaround to pass gtk_status_icon_position_menu as the GtkMenuPositionFunc.
func (v *Menu) PopupAtStatusIcon(statusIcon *StatusIcon, button gdk.Button, activateTime uint32) {
	C.gotk_menu_popup_at_status_icon(v.native(), statusIcon.native(), C.guint(button), C.guint32(activateTime))
}

/*
 * GtkMisc
 */

// Misc is a representation of GTK's GtkMisc.
type Misc struct {
	Widget
}

// native returns a pointer to the underlying GtkMisc.
func (v *Misc) native() *C.GtkMisc {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMisc(p)
}

func marshalMisc(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMisc(obj), nil
}

func wrapMisc(obj *glib.Object) *Misc {
	if obj == nil {
		return nil
	}

	return &Misc{Widget{glib.InitiallyUnowned{obj}}}
}

/*
 * End deprecated since version 3.14
 */
