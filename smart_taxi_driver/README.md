# Smart Taxi Driver

## Problem Statement

You are given a tree-like representation of a city, where each node represents a location, and edges represent roads connecting locations. The following information is provided:

- `N`: The total number of locations in the city.
- `K`: The maximum distance the taxi driver can travel before they must stop working.
- `P`: An array of size `N`, where `P[i]` represents the parent node of location `i` in the tree. The root of the tree is location `1`, and its parent is `0`.
- `C`: An array of size `N`, where `C[i]` represents the distance or cost of traveling along the road connecting location `i` to its parent `P[i]`.
- `T`: An array of size `N`, where `T[i]` represents the destination location of the person currently at location `i`. If `T[i]` is `0`, it means there is no one at location `i` needing a ride.

The taxi driver starts at location `1` (the root of the tree) and follows this process:

1. If there's a person at the current location (`T[current location] > 0`):
   - Calculate the distance to the person's destination.
   - If the driver has enough remaining distance to travel to the destination, drop the person off there, and then return to the root location (`1`).
   - Increment the count of people served.

2. Explore other locations connected to the current location (excluding the parent location, to avoid going back).

The taxi driver cannot travel more than a total distance of `K`.

### Task

Write a function to determine the maximum number of people the taxi driver can transport to their destinations.

### Input

- `N`: An integer representing the number of locations.
- `K`: An integer representing the maximum travel distance for the taxi driver.
- `P`: An array representing the parent nodes of each location.
- `C`: An array representing the distances of the roads.
- `T`: An array representing the destinations of people at each location.

### Output

- An integer representing the maximum number of people the taxi driver can serve.

### Constraints

- $1 \leq N \leq 10^5$
- $1 \leq K \leq 10^9$
- $0 \leq P[i] < i \quad \text{for } 1 \leq i \leq N$
- $1 \leq C[i] \leq 10^4 \quad \text{for } 1 \leq i \leq N$
- $0 \leq T[i] \leq N \quad \text{for } 1 \leq i \leq N$

## Approach

- [Approach-1](./solution_1/README.md#approach) (moderately optimized)
