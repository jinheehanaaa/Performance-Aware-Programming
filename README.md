
<details> <!-- Pre-Course-->
    <summary>Intro</summary>


<details>
    <summary>001 - Intro</summary>

- Understand CPU performance
- Aware of CPU instructions
- Comparison between Python & C
- SIMD & AVX for reducing the # of instructions
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

1. Cycle Benchbark on addition function
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


<details>
    <summary>004 - Single Instruction, Multiple Data</summary>

# SIMD
- - Make multiple ADD instructions in 1 instruction
- - PADDD => Load 4 things & add them to accumulator (Think of vector)
- - Slot Size for 32bit value: SSE (4), AVX(8), AVX512(16)
- - Storage size (register) is based on bit width
- - Need to use intrinsic function

# Benchmark in C (SSE & AVX)
- SingleScalar: 0.85 adds/clock peak
- SingleSSE: 3.12 adds/clock peak
- SingleAVX: 7.04 adds/clock peak
  
# Benchmark 2 in C (Unrolled version of AVX)
- QuadScalarPtr: 1.94
- SingleAVX: 7.08
- DualAVX: 11.01
- QuadAVX: 13.38

# Comparison
- C: 0.8 adds/cycle
- QuadAVX: 13.38 adds/cycle (16X performance of C)

</details>

<details>
    <summary>005 - Caching</summary>

# Load (Read)
- Load => Input[0]
- What determines the performance of load? Cache!

# Cache
## Register File in CPU
- desinged to produce value exteremly quickly & feed them to instruction
- Holds few hundreds of values at most

## Load Behavior
- CPU goes to {#} to find input[0]
1. L1 (~3 cycles, 32KB)
2. L2 (~14 cycles, 256KB)
3. L3 (~80 cycles, ~8MB)
4. Main Memory (~Slowest cycles, ~16GB)

## Boundary b/w L2 & L3 (Usually)
- L1, L2 is per core => Each core has 32KB, 256KB for L1, L2
- L3 is shared memory

# Benchmark L1, L2, L3 & Main Memory
- QuadAVXPtr 4096: 13.21 (L1 cache)
- QuadAVXPtr 32768: 7.70 (L2 cache)
- QuadAVXPtr 262144: 4.4 (L3 cache)
- QuadAVXPtr 33554432: 1.44 (Main Memory)

</details>

</details>  <!-- Pre-Course END-->


<details>
    <summary>Main Course</summary>
    
<details>
    <summary>The Haversine Distance Problem</summary>
    


</details>    


</details>





# Resources
- [Paid Tutorial](https://www.computerenhance.com/)



