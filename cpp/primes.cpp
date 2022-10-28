#include <cmath>
#include <iostream>
#include <vector>

int main(int count, char **args) {
  int n = 0;
  if (count < 3) {
    std::cout << "not enough arguments" << std::endl;
    return 0;
  }
  std::vector<int> nums;
  int start = std::stoi(args[1]);
  int end = std::stoi(args[2]);
  std::cout << "loading..." << std::endl;
  for (int n = start; n <= end; n++) {
    int i = 2;
    double n_root = sqrt(n);
    bool add = true;
    while (i <= n_root) {
      if (n % i == 0) {
        add = false;
        break;
      }
      i++;
    }
    if (add) {
      nums.push_back(n);
    }
  }

  std::cout << "primes from " << start << " to " << end << std::endl;
  for (int i = 0; i < nums.size(); i++) {
    std::cout << i + 1 << ") " << nums[i] << std::endl;
  }

  return 0;
}
