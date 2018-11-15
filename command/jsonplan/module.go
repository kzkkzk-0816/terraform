package jsonplan

// module is the representation of a module in state. This can be the root
// module or a child module.
type module struct {
	Resources []resource `json:"resources,omitempty"`

	// Address is the absolute module address, omitted for the root module
	Address string `json:"address,omitempty"`

	// Each module object can optionally have its own nested "child_modules",
	// recursively describing the full module tree.
	ChildModules []module `json:"child_modules,omitempty"`
}

type moduleCall struct {
	ResolvedSource    string                 `json:"resolved_source,omitempty"`
	Expressions       map[string]interface{} `json:"expressions,omitempty"`
	CountExpression   expression             `json:"count_expression,omitempty"`
	ForEachExpression expression             `json:"for_each_expression,omitempty"`
	Module            module                 `json:"module,omitempty"`
}
