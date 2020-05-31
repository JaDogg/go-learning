Go only things that start with capital are exported.
When two or more consecutive named function parameters share a type, you can omit the type from all but the last. 

---

 Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available. 


Return named

```go
func calc(x int, y int) (sum int, prod int) { 
    sum = x + y                               
    prod = x * y                              
    return                                    
}
```                                          


Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers


Switch without a condition is the same as switch true. 
This construct can be a clean way to write long if-then-else chains. 

```go 
t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("Good morning!")
case t.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}

```


Defer statement usage sample 


```go
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```



Arrays

```
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	fmt.Printf("Type of of names = %T\n", names) // Type = [4]string
	fmt.Printf("Type of a = %T\n", a)            // Type = []string

```


Maps


Mutating Maps
Insert or update an element in map m: 
```m[key] = elem```
Retrieve an element: 
```elem = m[key]```
Delete an element: 
```delete(m, key)```
Test that a key is present with a two-value assignment: 
```elem, ok = m[key]```
If key is in `m`, `ok` is `true`. If not, `ok` is `false`. 
If key is not in the map, then `elem` is the `zero value` for the map's element type. 
Note: If elem or ok have not yet been declared you could use a short declaration form: 
```elem, ok := m[key]```



Function values / lambda

```

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

```