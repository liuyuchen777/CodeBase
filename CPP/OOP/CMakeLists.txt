cmake_minimum_required(VERSION 3.0)

project(oop)

set(CMAKE_CXX_STANDARD 11)

include_directories(${CMAKE_CURRENT_SOURCE_DIR})

add_executable(basic basic.cpp warrior.h animal.h)

add_executable(advance advance.cpp)

add_executable(interview interview.cpp)