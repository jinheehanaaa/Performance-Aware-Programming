# What is haversine distance problem

The Haversine distance problem is a mathematical problem in which the distance between two points on a sphere (such as the Earth) is calculated using the Haversine formula. This formula takes into account the latitude and longitude of the two points, and computes the shortest distance along the surface of the sphere between them.

The Haversine formula is particularly useful in navigation and geolocation applications, where it is important to calculate distances between points on the Earth's surface with a high degree of accuracy. It is often used in GPS systems, mapping software, and other location-based services.

The formula is based on the law of haversines, which states that for a sphere of radius r, the haversine of a central angle between two points on the sphere is equal to the sum of the haversines of the latitudes and longitudes of the two points, minus the product of the two haversines.

The Haversine distance problem is a well-known problem in computer science, and efficient algorithms exist for its solution.

# Test in go
1. Haversine distance problem
2. Input file as JSON
3. Average (Ground truth)

# Terminology

Input: The time it took to read the input file, in seconds.
Math: The time it took to perform the haversine distance calculations, in seconds.
Total: The total time it took to run the program, in seconds (i.e., the sum of the input and math times).
Throughput: The number of haversine distances calculated per second, calculated as the total number of distances divided by the total time.
Note that the Throughput metric is calculated as the total number of haversine distances (i.e., len(points)-1) divided by the total time it took to read the input file and perform the calculations. This gives us an estimate of how many haversine distances our program can calculate per second. The Throughput metric can be useful for comparing the performance of different implementations of the haversine distance algorithm.