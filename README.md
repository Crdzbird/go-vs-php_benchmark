# BENCHMARK GO VS PHP

As you know a lot of business and enterprises decide to use PHP due to his long story of being "**simple**" and "**easy**" however more and more developers area searching for alternatives.
That's why i decided to create this repository.

*and being honest we as developers always search for the most clean/efficient and easy way to solve any problem.*

## How to run this project

To do the benchmark and see the results by your own eyes, please follow these steps.

 - git clone [https://github.com/Crdzbird/go-vs-php_benchmark](https://github.com/Crdzbird/go-vs-php_benchmark)
 - go to your go-vs-php_benchmark project location.
 - $ python3 generate-corpus.py -x 2000 -y 5000
 - $ php php/main.php 
 - $ go run go/main.go go/loader.go 
 - and finally wait to both programs finish their executions.
 you should probably see the time that requires each segment for both languages.

# Golang

**TL;DR**: [Why-You-Should-Learn-Go](https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65)

Go can be summarized in 5 main issues.

 1. **It Compiles Into Single Binary**

Golang built as a compiled language and Google developers did great job with it. Using _static linking_ it actually combining all dependency libraries and modules into one single binary file based on OS type and architecture. Which means if you are compiling your backend application on your laptop with Linux X86 CPU you can just upload compiled binary into server and it will work, without installing any dependencies there!

2. **Static Type System**

Type system is really important for large scale applications. Python is great and fun language but sometimes you are just getting unusual exceptions because trying to use variable as an integer but it turning out that it’s a string.

Go will let you know about this issue during compile time as a compiler error. This is where you winning time for this kind of stupid issues.

3.  **Performance!!**

Go performed better because of his concurrency model and CPU scalability. Whenever we need to process some internal request we are doing it with separate Goroutine, which are 10x cheaper in resources than Python 

> ( or any other programming Threads )

. So it save a lot of resources (Memory, CPU) because of the built in language features.

4. **You Don’t Need Web Framework For Go**

This is the most awesome thing about programming language. Go language creators and the community have built in so many tools natively supported by language core, that in most of the cases you really don’t need any 3rd party library. For example it has `http, json, html templating` built in language natively and you can build very complex API services without even thinking about finding library on Github!

But of course there is a lot of libraries and frameworks built for Go and making web applications with Go, but I’ll recommend build your web application or API service without any 3rd party library, because in most cases they are not making your life easier than using native packages.

5. **Great IDE support and debugging**

IDE support is one of the most important things when you are trying to switch your programming language. Comfortable IDE in average can save up to 80% of your coding time. I found [Go Plugin For JetBrains IDEA](https://github.com/go-lang-plugin-org/go-lang-idea-plugin) which has support also for (Webstorm, PHPStorm, etc…). That plugin is giving everything that you need for project development with the power of JetBrains IDEA you can really boost your development.
