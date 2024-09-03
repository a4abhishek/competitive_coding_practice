# Smart Taxi Driver - 2

## Problem Statement

You are given a city represented as a tree, where `N` people live, each at a unique place from 1 to `N`. Each person wants to go to a unique destination, ensuring that no two people want to go to the same destination. 

The road map of the city is given by an array `P` where each `P[i]` represents the parent node of the place `i`, and the distance between place `i` and its parent `P[i]` is given by an array `C`. The root of the tree is node `1`, meaning that `P[1] = 0` and `C[1] = 0`.

The city’s road map is a tree, and you, the taxi driver, need to transport as many people as possible under the following conditions:

1. The taxi driver can start from any place `x` in the city.
2. From any starting place `x`, the driver picks up the person at that place and drives them to their destination `T[x]`.
3. Upon reaching a destination, the driver immediately picks up the person there and drives them to their respective destination.
4. The taxi driver continues this process until they reach a place where no one is waiting.
5. The taxi driver cannot drive more than a total distance of `K`.

Your task is to find the **maximum number of people** the taxi driver can transport.

### Input

- `N`: An integer representing the number of people/places in the city.
- `K`: An integer representing the maximum travel distance for the taxi driver.
- `P`: An array representing the parent nodes of each place.
- `C`: An array representing the distances of the roads between places.
- `T`: An array representing the destination place for each person.

### Output

- An integer representing the maximum number of people the taxi driver can transport.

### Constraints

- $1 \leq N \leq 10^5$
- $1 \leq K \leq 10^9$
- $1 \leq C[i] \leq 10^4$ for $1 \leq i \leq N$
- $1 \leq T[i] \leq N$ and all values in `T` are unique permutations of `N$
