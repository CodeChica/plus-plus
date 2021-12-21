# Ruby

## irb

Irb stands for "Interactive Ruby".
The irb will read Ruby code that we've written, and return the result to us.
Let's try it out!

1. In your terminal, type `irb`, and press enter.
1. Type `name = "Your name"`. Replace with your name. Hit enter.
1. Now type `name` and enter. Your name should print to the screen!

```irb
$ irb
irb(main):001:0> puts "Hello, Ruby!"
Hello, Ruby!
=> nil
irb(main):002:0> name = "Your name"
=> "Your name"
irb(main):003:0> name
=> "Your name"
irb(main):004:0>
```

You've successfully used the irb! Let's see what else we can do with Ruby.

## Data types

### Integers

* integers (1, 10, 30000)
* floats (1.1232, -1.234324)

### Strings

Can be written with double (" ") or single (' ') quotes

```ruby
message = 'This is a string'
other_message = "So is this"
```

### Arrays - groups of things

```ruby
numbers = [1, 2, 3, 4, 5]
```

### Objects

* Objects hold data.
* Objects have "behavior".

Let's say we want to create several objects that are plants. We can start by creating a plant "class".

```ruby
class Plant
end
```

This class will allow us to create plant objects that hold their own data and exhibit behavior.
We can start by initializing a plant object to include some universal data that applies to any plant.

```ruby
class Plant
  attr_reader :color

  def initialize
    @color = "green"
  end
end
```

We can also allow a user to submit their own info when they create a plant object.

```ruby
class Plant
  attr_reader :color, :light, :height, :age

  def initialize(light, height, age)
    @color = "green"
    @light = light
    @height = height
    @age = age
  end
end
```

Now different plants can have their own light requirements, height, and age!
Finally, we can get our plant to do certain behaviors.

```ruby
class Plant
  attr_reader :color, :light, :height, :age

  def initialize(light, height, age)
    @color = "green"
    @light = light
    @height = height
    @age = age
  end

  def grow(sun)
    if sun == light
      height += 1
      puts "Thanks for the nutrients!"
    elsif sun > light
     puts "That's too much light!"
    elsif sun < light
      height -= 1
      "Help...need...sunlight"
    end
  end
end
```

Now a plant object is able to grow based on the amount of sunlight passed into the `grow` method.

### Resources

* [Intro to Ruby Cheatsheet](https://www.codecademy.com/learn/learn-ruby/modules/learn-ruby-introduction-to-ruby-u/cheatsheet)
