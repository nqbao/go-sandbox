# Learning Go

This is just a sandbox repository for me to put in random programming, mostly for me to learn go. The motivation is that I want to learn a new backend language to be able to write high-performance backend services, and there are many projects (k8s, docker, normad) are written in go. If i understand Go it will be easier for me to understand these projects.


What i like about Go:

  * Clear and nice syntax.
  * Channel and goroutine is awesome, a lot more cleaner than in python where i have to use Condition for cordination.
  * Concurrency is easy to reason. Everything is blocking, unless you let it `go`.
  * Strong-typed

What I don't like about Go:

  * Single workspace is a confusing.
  * Error handling is just awful. It's too bad that Go does not have a try / catch block.
  * JSON deserialization is hard if you don't know the key and type of the value.
  * Lacking of Generic

## Channels

It’s okie to leave channel open. GC will collect it

  * http://www.tapirgames.com/blog/golang-channel-closing

## Links

  * Package name convention
    * https://rakyll.org/style-packages
    * https://blog.golang.org/package-names
  * [Dependencies management with Dep](https://golang.github.io/dep)
    * [Should I commit vendor folder?](https://github.com/golang/dep/blob/master/docs/FAQ.md#should-i-commit-my-vendor-directory)
  * [appliedgo.net](https://appliedgo.net/tui/)
