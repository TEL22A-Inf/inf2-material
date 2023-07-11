package copybeispiel

func ExampleCopyDemoNaiv() {
	CopyDemoNaiv()

	// Output:
	// [1 2 3 4 5]
	// [1 2 3 4 5]
	// [55555 2 3 4 5]
	// [55555 2 3 4 5]
}

func ExampleCopyDemoCopy() {
	CopyDemoCopy()

	// Output:
	// [1 2 3 4 5]
	// [0 0 0 0 0]
	// [1 55555 3 4 5]
	// [1 2 3 4 5]
}
