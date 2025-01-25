# Go Array Index Out of Bounds Error

This repository demonstrates a common error in Go: exceeding the bounds of an array index.  The `bug.go` file contains the erroneous code, while `bugSolution.go` provides the corrected version.

The error occurs because Go arrays are zero-indexed and have a fixed size.  Attempting to access an element beyond the last valid index (length - 1) results in a runtime panic.

This example highlights the importance of careful index handling when working with arrays in Go.  Always ensure your loop conditions and index calculations respect the array's boundaries.