package erratum

func Use(o ResourceOpener, input string) (err error) {
	var r Resource
	for {
		r, err = o()
		if err == nil {
			break
		} else if _, ok := err.(TransientError); !ok {
			return err
		}
	}
	defer r.Close()
	defer func() {
		if x := recover(); x != nil {
			if frobErr, ok := x.(FrobError); ok {
				r.Defrob(frobErr.defrobTag)
			}
			err = x.(error)
		}
	}()
	r.Frob(input)
	return nil
}