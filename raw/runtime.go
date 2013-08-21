// Originally cloned from `github.com/tones111/raw`
//
// Copyright (c) 2011 Eleanor McHugh
//
// For addtional deatils, see the included `LICENSE` file

package raw

func Throw() {
	panic(nil)
}

func Catch(f func()) {
	defer func() {
		if x := recover(); x != nil {
			panic(x)
		}
	}()
	f()
}

func CatchAll(f func()) {
	defer func() {
		recover()
	}()
	f()
}
