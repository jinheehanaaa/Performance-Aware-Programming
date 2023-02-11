<details>
    <summary>001 - Intro</summary>
    - Understand CPU performance
    - Aware of CPU instructions
</details>


<details>
    <summary>002 - Waste</summary>
    
1. Waste Examples
## Definition
- ADD Instruction in assembly language & A+ = B
- LEA Instruction (Load Effective Address) & C=A+B
## Assembly Language
- ADD => 1 instruction
- LEA => 1 instruction
## C Language
- 1 instruction
## Python
- More than 100 instructions

2. Cycle Benchbark on addition function
## Python
- adds/cycle: 0.0061
## C Language
- adds/cycle: 0.80
- ~129x faster than python
- It can be more than ~1000x faster than python depending on the type of instruction

</details>

<details>
    <summary>003 - Instructions Per Clock</summary>

# IPC & ILP
- IPC (Instructions Per Clock) : average number of actual instruction that cpu executes on every clock cycle
- ILP (Instruction-Level Parallelism): # of instructions in cpu executing in cycle

# Typical For Loop Execution
```c
for(i = 0; i < count; i+=1)
{
    sum += input[i]
}
```
1. "add" inside loop
2. Get value ("load") from emmory
3. "add" for counter
2. "comparison"

# Benchmark
## "Unrolling" a loop
### Bnchmark: Unrolling Basic (PYTHON)
- Index +=2 & 2 Adds (Unroll2Scalar: 0.99 adds/closk peak)
- Index +=4 & 4 Adds (Unroll4Scalar: 0.99 adds/closk peak)
- It only helps for Index +=2 case
- PROBLEM: Serial Dependency Chain (every single ADD is dependent on previous ADD)
- We need to break the dependency chain among all the Add Inputs
- SOLUTION: Pairs of ADD (ex: 2 dependency chain)

### Benchmark: 2 dependency chain (PYTHON)
- Unroll2Scalar: 0.99
- DualScalar: 1.26
- QuadScalar: 1.70
- QuadScalarPtr: 1.94 (2X performance!)
- We can increase the performance of CPU by breaking the dependency chain
- 2X is samll compared to what we can do to improve this loop
</details>



# Resources
- [Paid Tutorial](https://www.computerenhance.com/)



