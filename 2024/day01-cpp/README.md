
# Dependencies

* cmake
* gtest 1.14.0 
* Ninja (optional)

# Generate

```
$ cmake -B Build -G Ninja
```

# Generate Debug

```
$ cmake -B Build -G Ninja -DCMAKE_BUILD_TYPE=Debug 
```

# Build

```
$ cmake --build Build 
```

# Run tests

```
$./Build/aoc-2024-day01-cpp
```
