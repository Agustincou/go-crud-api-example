package validators

type validator[T any] interface {
	LinkWith(validator[T]) validator[T]
	Check(T) error

	CheckNext(T) error
}

type Chain[T any] struct {
	nextValidator validator[T]
}

func (c *Chain[T]) CheckNext(req T) error {
	if c.nextValidator != nil {
		return c.nextValidator.Check(req)
	}

	return nil
}

func (c *Chain[T]) LinkWith(nextValidator validator[T]) validator[T] {
	c.nextValidator = nextValidator

	return nextValidator
}
