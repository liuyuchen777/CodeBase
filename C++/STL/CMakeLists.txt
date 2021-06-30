cmake_minimum_required (VERSION 2.8)

project(stl_demo)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -g -std=c++11")
# set(CMAKE_C_COMPILER "/usr/bin/gcc")
# set(CMAKE_CXX_COMPILER "/usr/bin/clang")

find_package(GTest REQUIRED) # find package of google test
find_package(Threads REQUIRED) # find package of multiple thread test

include_directories(${GTEST_INCLUDE_DIRS}) # add gtest dir

add_executable(${PROJECT_NAME} ) # add executable file

target_link_libraries(${PROJECT_NAME} ${GTEST_BOTH_LIBRARIES}) # add link library
target_link_libraries(${PROJECT_NAME} ${CMAKE_THREAD_LIBS_INIT})
add_test(Test ${PROJECT_NAME}) # add test
enable_testing()