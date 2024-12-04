#include <cstdlib>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>

bool report_safe(std::vector<int>);
bool report_safe_with_dampener(std::vector<int>);
bool same_polarity(int, int);
bool safe_difference(int);
std::vector<int> copy_report(std::vector<int>, int);

int main() {
	std::string file_path = "data/2024/day2.txt";

	std::ifstream file(file_path);
	if (!file.is_open()) {
		std::cerr << "Error: could not open the file: " << file_path << std::endl;
		return 1;
	}

	std::string line;
	std::vector<std::vector<int>> reports;

	while (std::getline(file, line)) {
		std::istringstream line_stream(line);
		std::string word;
		std::vector<int> report;
		while (line_stream >> word) {
			size_t idx;
			int i = std::stoi(word, &idx);
			if (idx == word.size()) {
				report.push_back(i);
			}
		}
		reports.push_back(report);
	}

	int safe = 0;
	int safe_with_dampener = 0;
	for (int i = 0; i < reports.size(); i++) {
		if (report_safe(reports[i])) {
			safe++;
			continue;
		}
		if (report_safe_with_dampener(reports[i])) {
			safe_with_dampener++;
		}
	}
	std::cout << "Safe reports: " << safe << std::endl;
	std::cout << "Safe reports with dampener: " << safe + safe_with_dampener << std::endl;

	return 0;
}

bool report_safe_with_dampener(std::vector<int> report) {
	for (int i = 0; i < report.size(); i++) {
		std::vector<int> temp = copy_report(report, i);
		if (report_safe(temp)) {
			return true;
		}
	}
	return false;
}

std::vector<int> copy_report(std::vector<int> report, int skip_idx) {
	std::vector<int> cp;
	for (int i = 0; i < report.size(); i++) {
		if (i == skip_idx) {
			continue;
		}
		cp.push_back(report[i]);
	}
	return cp;
}

bool report_safe(std::vector<int> report) {
	if (report.size() < 2) {
		return true;
	}

	int diff = report[0] - report[1];
	if (!safe_difference(diff)) {
		return false;
	}

	for (int i = 1; i < report.size()-1; i++) {
		int next = report[i] - report[i+1];
		if (!safe_difference(next)) {
			return false;
		}
		if (!same_polarity(diff, next)) {
			return false;
		}
		diff = next;
	}

	return true;
}

bool same_polarity(int a, int b) {
	if (a < 0 && b > 0) {
		return false;
	}
	if (a > 0 && b < 0) {
		return false;
	}
	return true;
}

bool safe_difference(int num) {
	if (num == 0) {
		return false;
	}
	if (std::abs(num) > 3) {
		return false;
	}
	return true;
}
