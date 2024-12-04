#include <cstdlib>
#include <fstream>
#include <iostream>
#include <string>

int main() {
	std::string file_path = "data/2024/day2.txt";

	std::ifstream file(file_path);
	if (!file.is_open()) {
		std::cerr << "Error: could not open the file: " << file_path << std::endl;
		return 1;
	}

	std::string report;
	while (std::getline(file, report)) {
		std::cout << report << std::endl;
	}

	return 0;
}

