#include <algorithm>
#include <iostream>
#include <fstream>
#include <unordered_map>
#include <vector>
#include <sstream>
#include <cstdlib>

int main() {
    std::string file_path = "data/2024/day1.txt";

    std::ifstream file(file_path);
    if (!file.is_open()) {
        std::cerr << "Error: Could not open the file: " << file_path << std::endl;
        return 1;
    }

    std::string line;
    std::vector<int> reports_a;
    std::vector<int> reports_b;
    while (std::getline(file, line)) {
		std::istringstream line_stream(line);
        std::string word;
		int list = 0;
        while (line_stream >> word) {
            try {
                size_t idx;
                int num = std::stoi(word, &idx);
                if (idx == word.size()) { 
					if (list == 0) {
						reports_a.push_back(num);
					} else if (list == 1) {
						reports_b.push_back(num);
					}
                }
            } catch (const std::invalid_argument& ignore) {
            } catch (const std::out_of_range& ignore) {
            }
			list++;
        }
    }

    file.close(); 
	
	std::sort(reports_a.begin(), reports_a.end());
	std::sort(reports_b.begin(), reports_b.end());

	int distance = 0;
	for (int i = 0; i < reports_a.size(); i++) {
		distance += std::abs(reports_a[i] - reports_b[i]);
	}

	std::unordered_map<int, int> cache;
	int similarity = 0;
	for (int i = 0; i < reports_a.size(); i++) {
		int key = reports_a[i];
		if (!cache.count(key)) {
			int count = 0;
			for (int j = 0; j < reports_b.size(); j++) {
				if (reports_b[j] == key) {
					count++;
				}
			}
			cache[key] = key * count;
		}
		similarity += cache[key];
	}

    std::cout << "Distance between lists: " << distance << std::endl;
    std::cout << "Similarity between lists: " << similarity << std::endl;

    return 0;
}
