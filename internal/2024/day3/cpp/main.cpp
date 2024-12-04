#include <algorithm>
#include <cstring>
#include <fstream>
#include <iostream>
#include <regex>
#include <stdexcept>
#include <string>
#include <variant>

int sum_products(std::string);
int sum_products_conditional(std::string);
int multiply(std::string);

int main() {
	std::string file_path = "data/2024/day3.txt";

	std::ifstream file(file_path);
	if (!file.is_open()) {
		std::cerr << "Error: could not open file " << file_path << std::endl;
		return 1;
	}

	std::string line;
	std::string full;

	while (std::getline(file, line)) {
		full += line;
	}
	
	int sum = sum_products(full);
	int sum_conditional = sum_products_conditional(full);

	std::cout << "Sum of products: " << sum << std::endl;
	std::cout << "Sum of conditional products: " << sum_conditional << std::endl;

	return 0;
}

int sum_products_conditional(std::string line) {
	int sum = 0;
	size_t start = 0;

	while (start != std::variant_npos) {
		size_t end = std::min(line.find("don't()", start), line.length());	
		sum += sum_products(line.substr(start, end-start));
		start = line.find("do()", end);
	}
	return sum;
}

int sum_products(std::string line) {
	std::regex re("mul\\((\\d*,\\d*)\\)");
	std::smatch match; 

	int sum = 0;
	while (std::regex_search(line, match, re)) {
		sum += multiply(match[1]);
		line = match.suffix().str();
	}

	return sum;
}

int multiply(std::string nums_as_str) {
	size_t c = nums_as_str.find(",");
	if (c == std::variant_npos) {
		return 0;
	}

	try {
		size_t idx;

		std::string a = nums_as_str.substr(0, c);
		int a_num = std::stoi(a, &idx);
		if (idx != a.length()) {
			return 0;
		}

		std::string b = nums_as_str.substr(c+1, nums_as_str.length());
		int b_num = std::stoi(b, &idx);
		if (idx != b.length()) {
			return 0;
		}
		return a_num * b_num;

	} catch(std::invalid_argument ignore) {}

	return 0;
}
