func myFunc(x int) (int, error) {
  defer func() {
    if r := recover(); r != nil {
      // Handle the panic and return an error
      // You can log the panic for debugging purposes here

      // return an approriate error for the situation
      return 0, fmt.Errorf("panic occurred: %v", r)
    }
  }()
  if x == 0 {
    panic("cannot divide by zero")
  }
  return 10 / x, nil
}