cmake_minimum_required(VERSION 3.8)
project(aoc-2024-day01-cpp)

set(GOOGLETEST_VERSION 1.14.0)
find_package(GTest ${GOOGLETEST_VERSION}) 

file(GLOB SOURCES "*.cpp")
add_executable(${PROJECT_NAME} ${SOURCES})
target_link_libraries(${PROJECT_NAME} gtest)
