#include <fstream>
#include <iostream>
#include <string>
#include <vector>

int countXmasses(std::vector<std::vector<char>>, int, int);
bool hasXMas(std::vector<std::vector<char>>);

const int cardinalities[8][2] = {{-1, -1}, {-1,0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}};

int main() {
	std::string file_path = "data/2024/day4.txt";

	std::ifstream file(file_path);
	if (!file.is_open()) {
		std::cerr << "Error: could not open file " << file_path << std::endl;
		return 1;
	}

	std::string line;
	std::vector<std::vector<char>> data;
	while (std::getline(file, line)) {
		std::vector<char> row;
		for (int i = 0; i < line.length(); i++) {
			row.push_back(line[i]);
		}
		data.push_back(row);
	}

	int xmasses = 0;
	for (int i = 0; i < data.size(); i++) {
		for (int j = 0; j < data[i].size(); j++) {
			xmasses += countXmasses(data, i, j);
		}
	}
	std::cout << "XMAS count: " << xmasses << std::endl;

	int xMasses = 0;
	for (int i = 0; i < data.size()-2; i++) {
		for (int j = 0; j < data[i].size()-2; j++) {
			std::vector<std::vector<char>> temp;
			for (int r = 0; r < 3; r++) {
				std::vector<char> temp_row = {data[i+r][j], data[i+r][j+1], data[i+r][j+2]};
				temp.push_back(temp_row);
			}
			if (hasXMas(temp)) {
				xMasses++;
			}
		}
	}

	std::cout << "X-MAS count: " << xMasses << std::endl;
	return 0;
}

bool hasXMas(std::vector<std::vector<char>> data) {
	if (data[1][1] != 'A') {
		return false;
	}
	std::string target = "MS";

	std::string tr = std::string(1, data[2][2]) + std::string(1, data[0][0]);
	std::string rt = std::string(1, data[0][0]) + std::string(1, data[2][2]);
	if (tr != target && rt != target) {
		return false;
	}

	std::string br = std::string(1, data[2][0]) + std::string(1, data[0][2]);
	std::string rb = std::string(1, data[0][2]) + std::string(1, data[2][0]);
	return br == target || rb == target;
}

int countXmasses(std::vector<std::vector<char>> data, int y, int x) {
	int count = 0;
	std::string target = "XMAS";
	bool found = false;
	for (auto card:cardinalities) {
		found = true;
		for (int i = 0; i < target.length(); i++) {
			int posY = y + card[0] * i;
			int posX = x + card[1] * i;
			if (posY < 0 || posY >= data.size()) {
				found = false;
				break;
			}
			if (posX < 0 || posX >= data[posY].size()) {
				found = false;
				break;
			}
			if (data[posY][posX] != target[i]) {
				found = false;
				break;
			}
		}
		if (found) {
			count++;
		}
	}
	return count;
}
