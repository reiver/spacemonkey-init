package main

func do() error {

	if err := doRook(); nil != err {
		return err
	}

	if err := doGemini(); nil != err {
		return err
	}

	return nil
}
