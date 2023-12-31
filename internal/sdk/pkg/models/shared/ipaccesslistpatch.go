// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type IPAccessListPatch struct {
	// Elements to add. Executed after "remove" part is processed.
	Add []IPAccessListEntry `json:"add,omitempty"`
	// Elements to remove. Executed before "add" part is processed.
	Remove []IPAccessListEntry `json:"remove,omitempty"`
}
