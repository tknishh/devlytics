# Problem 2
Write Go code to perform data delivery from one goroutine, and processing of the data in another goroutine¶

## Explanation
Start three goroutines, call them g1, g2, g3
Pass an integer to goroutine g1, and g2
g1,g2 will send this data directly into g3
g3 adds the integers

## Note
You may not need any libraries, because this assignment makes use of just the default stuff Go programming language provides to you. Through this assignment you'll learn how two goroutines communicate with each other
