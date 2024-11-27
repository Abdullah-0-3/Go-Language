# Go Language

Go is quickly becoming one of the top programming languages. While it’s somewhat similar to Python, Python is still more versatile for many tasks. 

However, Go stands out with several advantages that make it a popular choice for developers.

<!-- ![Go Language](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRtcufyYKf3mr5OIHmufcWXbUw1WPXw4nRtSg&s) -->

<img src="https://www.nixsolutions.com/uploads/2020/07/Golang.png" alt="Go Language" style="width:100%; height:100%;">

### Why Choose Go?

- **Lightweight:** Go is simple and efficient.
- **Blazing Fast:** It’s much faster compared to many other languages.
- **High Performance:** Ideal for building high-speed and scalable applications.

---
I’ve been exploring Go with **ChatGPT**, and it’s been a game-changer for improving my coding skills and development logic. If you haven’t tried it yet, you should! This is just the beginning.

---
### Task 01: **Hello World Program**
Write a Go program that prints "Hello, World!" to the console.

---
### Task 02: **Arthimetic Opreators**
Write a Go program that takes two integers as input from the user and outputs their sum, difference, product, and quotient.

---
### Task 03: **FizzBuzz**
Write a Go program that prints the numbers from 1 to 100. For multiples of three, print "Fizz" instead of the number, and for the multiples of five, print "Buzz". For numbers that are multiples of both three and five, print "FizzBuzz".

---
### Task 04: **Count Vowels and Consonants**
Write a Go program that takes a string input from the user and counts the number of vowels and consonants in that string. For this task, consider the English vowels to be A, E, I, O, and U (both uppercase and lowercase).

Your program should ignore any non-alphabetic characters. Here’s how you might approach this:

1. Read input from the user.
2. Initialize counters for vowels and consonants.
3. Loop through each character in the string to determine if it is a vowel, consonant, or neither.
4. Print the results.

---
### Task 05: **Simple Calculator**
Write a Go program that functions as a simple calculator. Your program should:

Prompt the user to enter two numbers.
Prompt the user to choose an operation (addition, subtraction, multiplication, or division).

Perform the chosen operation on the numbers and print the result.
Here’s a quick guide to get you started:

- Use a loop to allow multiple calculations until the user decides to exit.
- Handle division by zero in the case of division operation.
- Use functions for different operations to keep your code organized.

---
### Task 06: **Fibonacci Series**
Write a Go program that generates the Fibonacci series up to a user-specified number of terms. 

The Fibonacci series starts with 0 and 1, and then each subsequent number is the sum of the previous two. For example, the series looks like this:
0, 1, 1, 2, 3, 5, 8, 13, ...

Your program should:

- Prompt the user for the number of terms they want in the Fibonacci series.
- Use a loop to calculate and print the Fibonacci numbers up to that number of terms.
- Handle cases where the input is less than or equal to zero (for example, you can print a message indicating that the number of terms must be positive).

--- 
### Task 07: **Basic File I/O**
Write a Go program that reads a text file and counts the number of words in it. 

**Requirements:**

- Prompt the user to enter the filename (make sure the file exists in the same directory as your Go program).
- Open the file and read its contents.
- Count the number of words in the file (consider words to be separated by spaces).
Print the total word count to the console.

**Extras:**

- If the file does not exist, handle the error gracefully and inform the user.
- You can also consider counting lines or characters for additional functionality.

---
### Task 08: **Simple HTTP Server**
Write a simple HTTP server in Go that responds with "Hello, World!" when accessed.

**Requirements:**

- Create a basic HTTP server that listens on a specified port (for example, port 8080).
- When the server receives a GET request on the root path (/), it should respond with "Hello, World!".
- You can test your server by opening a web browser or using tools like curl to access http://localhost:8080.

**Extras:**

- Handle any potential errors that come from starting the server.
- Extend the server to handle additional paths (e.g., /goodbye) and respond with different messages.

---
### Task 09: **JSON Serialization and Deserialization**
Create a Go program that defines a simple struct, serializes it to JSON, and then deserializes it back to a struct.

**Requirements:**

1. Define a struct that represents a basic entity (for example, a Person with fields like Name, Age, and Email).
2. Create an instance of that struct, populate it with some data, and serialize it to JSON format.
3. Print the resulting JSON string to the console.
4. Deserialize the JSON string back to a struct, and print the resulting struct data to the console.

**Extras:**

- Handle any potential errors during serialization and deserialization.
- Allow for the input of the struct data from the user before serialization.