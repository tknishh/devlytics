# Problem 1
Build a webserver with a couple of endpoints

## Explanation:
A webserver is a system that can receive a request, process it, and return a response. We'll have two endpoints, let's call it endpoint-one, and endpoint-two.

To connect to an endpoint, we'll have to do something like

`curl ip-address:port/endpoint-one`

- endpoint-one doesn't require an input data, but returns a response one.
- endpoint-two takes an input data like a string, and returns a response string-two, where string could be anything

## Frameworks
Programming Language : Golang, and related libraries
Use gomux for the webserver

