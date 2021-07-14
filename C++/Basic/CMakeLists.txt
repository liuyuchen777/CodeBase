cmake_minimum_required(VERSION 3.0)

project(basic_test)

set(CMAKE_CXX_STANDARD 11)

include_directories(${CMAKE_CURRENT_SOURCE_DIR})

add_executable(hello_world hello_world.cpp)