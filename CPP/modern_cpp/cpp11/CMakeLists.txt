cmake_minimum_required(VERSION 3.0)

project(c++11_test)

set(CMAKE_CXX_STANDARD 11)

include_directories(${CMAKE_CURRENT_SOURCE_DIR})

add_executable(type_deduce type_deduce.cc c++11_test.h)

add_executable(using_alias using_alias.cc c++11_test.h)

add_executable(template_default_type template_default_type.cpp)

add_executable(tuple tuple.cpp)

add_executable(list_initialize list_initialize.cpp)

add_executable(lambda_fun lambda_fun.cpp)

add_executable(new_data_type new_data_type.cpp)

add_executable(constexpr constexpr.cpp)

add_executable(cpp_reference cpp_reference.cpp)

add_executable(smart_pointer smart_pointer.cpp)