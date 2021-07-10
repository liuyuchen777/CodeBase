cmake_minimum_required(VERSION 2.8)

project(c++11_test)

set(CMAKE_CXX_STANDARD 11)

include_directories(${CMAKE_CURRENT_SOURCE_DIR})

add_executable(type_deduce type_deduce.cc c++11_test.h)

add_executable(using_alias using_alias.cc c++11_test.h)