# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.19

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Disable VCS-based implicit rules.
% : %,v


# Disable VCS-based implicit rules.
% : RCS/%


# Disable VCS-based implicit rules.
% : RCS/%,v


# Disable VCS-based implicit rules.
% : SCCS/s.%


# Disable VCS-based implicit rules.
% : s.%


.SUFFIXES: .hpux_make_needs_suffix_list


# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = "/Users/yuchenliu/Library/Application Support/JetBrains/Toolbox/apps/CLion/ch-0/211.7628.27/CLion.app/Contents/bin/cmake/mac/bin/cmake"

# The command to remove a file.
RM = "/Users/yuchenliu/Library/Application Support/JetBrains/Toolbox/apps/CLion/ch-0/211.7628.27/CLion.app/Contents/bin/cmake/mac/bin/cmake" -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /Users/yuchenliu/Local_Lab/CodeBase/C++

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug

# Include any dependencies generated for this target.
include Basic/CMakeFiles/basic_test.dir/depend.make

# Include the progress variables for this target.
include Basic/CMakeFiles/basic_test.dir/progress.make

# Include the compile flags for this target's objects.
include Basic/CMakeFiles/basic_test.dir/flags.make

Basic/CMakeFiles/basic_test.dir/hello_world.cpp.o: Basic/CMakeFiles/basic_test.dir/flags.make
Basic/CMakeFiles/basic_test.dir/hello_world.cpp.o: ../Basic/hello_world.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object Basic/CMakeFiles/basic_test.dir/hello_world.cpp.o"
	cd /Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug/Basic && /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/basic_test.dir/hello_world.cpp.o -c /Users/yuchenliu/Local_Lab/CodeBase/C++/Basic/hello_world.cpp

Basic/CMakeFiles/basic_test.dir/hello_world.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/basic_test.dir/hello_world.cpp.i"
	cd /Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug/Basic && /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /Users/yuchenliu/Local_Lab/CodeBase/C++/Basic/hello_world.cpp > CMakeFiles/basic_test.dir/hello_world.cpp.i

Basic/CMakeFiles/basic_test.dir/hello_world.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/basic_test.dir/hello_world.cpp.s"
	cd /Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug/Basic && /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /Users/yuchenliu/Local_Lab/CodeBase/C++/Basic/hello_world.cpp -o CMakeFiles/basic_test.dir/hello_world.cpp.s

# Object files for target basic_test
basic_test_OBJECTS = \
"CMakeFiles/basic_test.dir/hello_world.cpp.o"

# External object files for target basic_test
basic_test_EXTERNAL_OBJECTS =

Basic/basic_test: Basic/CMakeFiles/basic_test.dir/hello_world.cpp.o
Basic/basic_test: Basic/CMakeFiles/basic_test.dir/build.make
Basic/basic_test: Basic/CMakeFiles/basic_test.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking CXX executable basic_test"
	cd /Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug/Basic && $(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/basic_test.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
Basic/CMakeFiles/basic_test.dir/build: Basic/basic_test

.PHONY : Basic/CMakeFiles/basic_test.dir/build

Basic/CMakeFiles/basic_test.dir/clean:
	cd /Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug/Basic && $(CMAKE_COMMAND) -P CMakeFiles/basic_test.dir/cmake_clean.cmake
.PHONY : Basic/CMakeFiles/basic_test.dir/clean

Basic/CMakeFiles/basic_test.dir/depend:
	cd /Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /Users/yuchenliu/Local_Lab/CodeBase/C++ /Users/yuchenliu/Local_Lab/CodeBase/C++/Basic /Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug /Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug/Basic /Users/yuchenliu/Local_Lab/CodeBase/C++/cmake-build-debug/Basic/CMakeFiles/basic_test.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : Basic/CMakeFiles/basic_test.dir/depend
