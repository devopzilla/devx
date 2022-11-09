package devx

#Application: {
	$guku: "Application"

	components: #Components
}

#Components: [Id=string]: #Component & {
	$guku: id: Id
}

#Component: {
	$guku: {
		component: string
		id:        string
		children?: _
	}
	#Trait
	...
}

#Trait: {
	$guku: traits: [string]: _
	...
}

#Transformer: {
	$guku: transformer: string

	input: {
		component: #Component
		context:   _
	}

	feedforward: {
		components: #Components
	}

	feedback: {
		component: input.component
	}
}

// func feedforward(context, component) -> components
// func feedback(components) -> component
