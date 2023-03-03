<details>
    <summary>002 - "Clean" Code, Horrible Performance</summary>

# Clean Code Definition 
1. Polymorphism & Get rid of if/switch statement 
2. No internals (Should always call polymorphic functions, but not look into the object)
3. Small function
4. Function should do 1 thing
5. Don't repeat yourself (D.R.Y)

# Let's break the rules
- breaking rules #1, #2 => 11X faster for our test
- breaking rules #3 => 15x faster for our test

</details>
