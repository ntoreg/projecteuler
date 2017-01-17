# This source was generated from grumpy.
# $ grumpy hazumi.py > hazumi_grumpy.go
# see: https://github.com/google/grumpy

package main
import (
	πg "grumpy"
	π_os "os"
)
func initModule(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
	ßTrue := πg.InternStr("True")
	ß__main__ := πg.InternStr("__main__")
	ß__name__ := πg.InternStr("__name__")
	ßa := πg.InternStr("a")
	ßi := πg.InternStr("i")
	ßis_hazumi := πg.InternStr("is_hazumi")
	ßlist := πg.InternStr("list")
	ßsorted := πg.InternStr("sorted")
	ßstr := πg.InternStr("str")
	ßx := πg.InternStr("x")
	var πTemp001 *πg.Object
	_ = πTemp001
	var πTemp002 []πg.FunctionArg
	_ = πTemp002
	var πTemp003 *πg.Object
	_ = πTemp003
	var πTemp004 *πg.Object
	_ = πTemp004
	var πTemp005 bool
	_ = πTemp005
	var πTemp006 *πg.Object
	_ = πTemp006
	var πTemp007 []*πg.Object
	_ = πTemp007
	var πTemp008 []*πg.Object
	_ = πTemp008
	var πTemp009 bool
	_ = πTemp009
	var πTemp010 *πg.Object
	_ = πTemp010
	var πTemp011 *πg.Object
	_ = πTemp011
	var πE *πg.BaseException; _ = πE
	for ; πF.State() >= 0; πF.PopCheckpoint() {
		switch πF.State() {
		case 0:
		default: panic("unexpected function state")
		}
		// line 5: def is_hazumi(a):
		πF.SetLineno(5)
		πTemp002 = make([]πg.FunctionArg, 1)
		πTemp002[0] = πg.FunctionArg{Name: "a", Def: nil}
		πTemp001 = πg.NewFunction(πg.NewCode("is_hazumi", "../hazumi.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
			var µa *πg.Object = πArgs[0]; _ = µa
			var πTemp001 *πg.Object
			_ = πTemp001
			var πTemp002 *πg.Object
			_ = πTemp002
			var πTemp003 bool
			_ = πTemp003
			var πTemp004 *πg.Object
			_ = πTemp004
			var πTemp005 []*πg.Object
			_ = πTemp005
			var πTemp006 *πg.Object
			_ = πTemp006
			var πTemp007 *πg.Object
			_ = πTemp007
			var πTemp008 πg.KWArgs
			_ = πTemp008
			var πE *πg.BaseException; _ = πE
			for ; πF.State() >= 0; πF.PopCheckpoint() {
				switch πF.State() {
				case 0:
				default: panic("unexpected function state")
				}
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp005 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp005[0] = µa
				if πTemp006, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
					continue
				}
				if πTemp007, πE = πTemp006.Call(πF, πTemp005, nil); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp005)
				if πTemp004, πE = πg.Eq(πF, µa, πTemp007); πE != nil {
					continue
				}
				πTemp002 = πTemp004
				if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
					continue
				}
				if πTemp003 {
					goto Label1
				}
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp005 = πF.MakeArgs(1)
				if πE = πg.CheckLocal(πF, µa, "a"); πE != nil {
					continue
				}
				πTemp005[0] = µa
				if πTemp006, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				πTemp008 = πg.KWArgs{
					{"reverse", πTemp006},
				}
				if πTemp006, πE = πg.ResolveGlobal(πF, ßsorted); πE != nil {
					continue
				}
				if πTemp007, πE = πTemp006.Call(πF, πTemp005, πTemp008); πE != nil {
					continue
				}
				πF.FreeArgs(πTemp005)
				if πTemp004, πE = πg.Eq(πF, µa, πTemp007); πE != nil {
					continue
				}
				πTemp002 = πTemp004
			Label1:
				if πTemp003, πE = πg.IsTrue(πF, πTemp002); πE != nil {
					continue
				}
				πTemp001 = πg.GetBool(!πTemp003).ToObject()
				if πTemp003, πE = πg.IsTrue(πF, πTemp001); πE != nil {
					return nil, πE
				}
				if πTemp003 {
					goto Label2
				}
				goto Label3
				// line 6: if not (a == sorted(a) or a == sorted(a, reverse=True)):
				πF.SetLineno(6)
			Label2:
				// line 7: return True
				πF.SetLineno(7)
				if πTemp001, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
					continue
				}
				return πTemp001, nil
				goto Label3
			Label3:
				return nil, nil
			}
			return nil, πE
		}), πF.Globals()).ToObject()
		if πE = πF.Globals().SetItem(πF, ßis_hazumi.ToObject(), πTemp001); πE != nil {
			continue
		}
		if πTemp004, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
			continue
		}
		if πTemp003, πE = πg.Eq(πF, πTemp004, ß__main__.ToObject()); πE != nil {
			continue
		}
		if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
			return nil, πE
		}
		if πTemp005 {
			goto Label1
		}
		goto Label2
		// line 10: if __name__ == "__main__":
		πF.SetLineno(10)
	Label1:
		// line 12: x = 0
		πF.SetLineno(12)
		if πE = πF.Globals().SetItem(πF, ßx.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
			continue
		}
		// line 13: i = 0
		πF.SetLineno(13)
		if πE = πF.Globals().SetItem(πF, ßi.ToObject(), πg.NewInt(0).ToObject()); πE != nil {
			continue
		}
		// line 14: while True:
		πF.SetLineno(14)
	Label3:
		if πTemp003, πE = πg.ResolveGlobal(πF, ßTrue); πE != nil {
			continue
		}
		if πTemp005, πE = πg.IsTrue(πF, πTemp003); πE != nil {
			continue
		}
		if !πTemp005 {
			goto Label4
		}
		// line 15: i += 1
		πF.SetLineno(15)
		if πTemp004, πE = πg.ResolveGlobal(πF, ßi); πE != nil {
			continue
		}
		if πTemp006, πE = πg.IAdd(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßi.ToObject(), πTemp006); πE != nil {
			continue
		}
		// line 16: a = list(str(i))
		πF.SetLineno(16)
		πTemp007 = πF.MakeArgs(1)
		πTemp008 = πF.MakeArgs(1)
		if πTemp004, πE = πg.ResolveGlobal(πF, ßi); πE != nil {
			continue
		}
		πTemp008[0] = πTemp004
		if πTemp004, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
			continue
		}
		if πTemp006, πE = πTemp004.Call(πF, πTemp008, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp008)
		πTemp007[0] = πTemp006
		if πTemp004, πE = πg.ResolveGlobal(πF, ßlist); πE != nil {
			continue
		}
		if πTemp006, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp007)
		if πE = πF.Globals().SetItem(πF, ßa.ToObject(), πTemp006); πE != nil {
			continue
		}
		πTemp007 = πF.MakeArgs(1)
		if πTemp004, πE = πg.ResolveGlobal(πF, ßa); πE != nil {
			continue
		}
		πTemp007[0] = πTemp004
		if πTemp004, πE = πg.ResolveGlobal(πF, ßis_hazumi); πE != nil {
			continue
		}
		if πTemp006, πE = πTemp004.Call(πF, πTemp007, nil); πE != nil {
			continue
		}
		πF.FreeArgs(πTemp007)
		if πTemp009, πE = πg.IsTrue(πF, πTemp006); πE != nil {
			return nil, πE
		}
		if πTemp009 {
			goto Label5
		}
		goto Label6
		// line 17: if is_hazumi(a):
		πF.SetLineno(17)
	Label5:
		// line 18: x += 1
		πF.SetLineno(18)
		if πTemp004, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
			continue
		}
		if πTemp006, πE = πg.IAdd(πF, πTemp004, πg.NewInt(1).ToObject()); πE != nil {
			continue
		}
		if πE = πF.Globals().SetItem(πF, ßx.ToObject(), πTemp006); πE != nil {
			continue
		}
		if πTemp011, πE = πg.ResolveGlobal(πF, ßx); πE != nil {
			continue
		}
		if πTemp010, πE = πg.Mul(πF, πTemp011, πg.NewFloat(1.0).ToObject()); πE != nil {
			continue
		}
		if πTemp011, πE = πg.ResolveGlobal(πF, ßi); πE != nil {
			continue
		}
		if πTemp006, πE = πg.Div(πF, πTemp010, πTemp011); πE != nil {
			continue
		}
		if πTemp004, πE = πg.Eq(πF, πTemp006, πg.NewFloat(0.99).ToObject()); πE != nil {
			continue
		}
		if πTemp009, πE = πg.IsTrue(πF, πTemp004); πE != nil {
			return nil, πE
		}
		if πTemp009 {
			goto Label7
		}
		goto Label8
		// line 19: if x * 1.0 / i == 0.99:
		πF.SetLineno(19)
	Label7:
		// line 20: print i
		πF.SetLineno(20)
		πTemp007 = make([]*πg.Object, 1)
		if πTemp004, πE = πg.ResolveGlobal(πF, ßi); πE != nil {
			continue
		}
		πTemp007[0] = πTemp004
		if πE = πg.Print(πF, πTemp007, true); πE != nil {
			continue
		}
		// line 21: break
		πF.SetLineno(21)
		goto Label4
		goto Label8
	Label8:
		goto Label6
	Label6:
		goto Label3
		goto Label4
	Label4:
		goto Label2
	Label2:
		// line 30: """
		πF.SetLineno(30)
		return nil, nil
	}
	return nil, πE
}
var Code *πg.Code
func main() {
	Code = πg.NewCode("<module>", "../hazumi.py", nil, 0, initModule)
	π_os.Exit(πg.RunMain(Code))
}

/*
$ time ./hazumi_grumpy
477300

real	0m5.787s
user	0m6.125s
sys	0m0.087s
*/
