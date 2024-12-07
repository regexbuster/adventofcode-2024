# Day 1 (Advent of Code 2024)

**_Disclaimer:_** _I will only be using the examples from the text as my system's input as requested by AOC's creator, Eric Wastl, when doing write ups of my work and thought process._

## Part 1's Challenge

Given two lists of numbers in any order, give the sum of the difference between each pair of values. Pairs of values are constructed using the Nth smallest value from each list. The problem gives these two lists as the example:

```
3   4
4   3
2   5
1   3
3   9
3   3
```

While not given in the challenge specifications I'm going to sort each list from smallest to largest value for clarity.

```
1   3
2   3
3   3
3   4
3   5
4   9
```

Then we can calculate the difference between each pair:

```
|1-3| = 2
|2-3| = 1
|3-3| = 0
|3-4| = 1
|3-5| = 2
|4-9| = 5
```

And then when we sum it up it becomes `11`!

```
2 + 1 + 0 + 1 + 2 + 5 = 11
```

## My Part 1 Solution

First I started by making a quick function that just panics if an err occurred so I don't have to write if statements to check for errors across my code. It's not super necessary to pull this out into it's own function but it does clean up future code and is simple to remember what it does when you see it in the code.

```go
func check(e error){
	if e != nil {
		panic(e)
	}
}
```

I then write some logic to open the puzzle input file, define two slices that will operate as my two lists, and then what loops through the file and adds the numbers to each list. I could have saved myself some future trouble by converting each string version of the numbers into ints now but I decided to just convert them when needed. _I will probably change that when I refine my answers in the future._ While it may be more verbose that needed, it was a simple matter of splitting each line and appending it to their list. I added them to try and solve a different issue and they managed to persist till now. _This is another thing I will probably change now that I've seen it works without the checks._

```go
file, err := os.Open("./puzzleInput.txt")
defer file.Close()

check(err)

s := bufio.NewScanner(file)

// holds the values for each list
var listOne []string
var listTwo []string

// scan through the lists and adds each side to their above slice
for s.Scan() {
	splitted := strings.Split(s.Text(), "   ")

	if splitted[0] != ""{
		listOne = append(listOne, splitted[0])
	}

	if splitted[1] != ""{
		listTwo = append(listTwo, splitted[1])
	}
}
```

I use the slices package to quickly sort each slice list:

```go
slices.Sort(listOne)
slices.Sort(listTwo)
```

Finally, I run through each list and sum up the differences between each Nth smallest pair of values:

```go
sum  :=  0

// convert strings to int and then add the difference abs(num1 - num2)
for  i  :=  0; i  <  len(listOne); i++ {
	g, errg  :=  strconv.Atoi(listOne[i])
	check(errg)

	h, errh  :=  strconv.Atoi(listTwo[i])
	check(errh)

	sum  += (abs(g-h))
}

fmt.Printf("Total Distance: %d\n", sum)
```

I also wrote a quick function to do absolute values instead of importing the math package.

```go
func  abs(i  int) int{
	if  i  >=  0{
		return  i
	}

	return  i  *  -1
}
```

There may be some small improvements I've missed from there but otherwise this is the finished product. I quick sorter and summer to find the total distance between all the Nth smallest pairs!

```
> go run part1.go
Total Distance: 11
```
