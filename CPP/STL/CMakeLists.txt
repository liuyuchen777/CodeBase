cmake_minimum_required (VERSION 2.8)

project(stl_demo)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -g -std=c++11")
set(CMAKE_CXX_COMPILER "/usr/bin/g++")

find_package(GTest REQUIRED) # find package of google test
find_package(Threads REQUIRED) # find package of multiple thread test

include_directories(${GTEST_INCLUDE_DIRS}) # add gtest dir

add_executable(${PROJECT_NAME} 
                test.cc 
                vector_demo.cc vector_demo.h
                set_demo.cc set_demo.h
                map_demo.cc map_demo.h
                unordered_map_demo.cc unordered_map_demo.h
                unordered_set_demo.cc unordered_set_demo.h
            ) # add executable file

target_link_libraries(${PROJECT_NAME} ${GTEST_BOTH_LIBRARIES}) # add link library
target_link_libraries(${PROJECT_NAME} ${CMAKE_THREAD_LIBS_INIT})
add_test(Test ${PROJECT_NAME}) # add test
enable_testing()