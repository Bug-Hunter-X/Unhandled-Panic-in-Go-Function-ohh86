# Unhandled Panic in Go Function

This repository demonstrates a common error in Go: unhandled panics in functions that should return errors.

The `bug.go` file shows a function that returns an error when dividing by zero, but this solution is incomplete because a panic will occur instead of the error being returned if x is 0.

The solution in `bugSolution.go` addresses this. This shows the corrected code that uses `recover` within a deferred function to handle potential panics and return a proper error to the caller.  This ensures predictable behavior even in exceptional cases.