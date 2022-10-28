import sys
import math

if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("not enough arguments")
        sys.exit()

    start = int(sys.argv[1])
    end = int(sys.argv[2])
    nums = []
    print("loading...")
    for n in range(start,end+1):
        n_root = math.sqrt(n)
        i = 2
        add = True
        while i <= n_root:
            if n % i == 0:
                add = False
                break
            i+=1 
    
        if add == True:
            nums.append(n)

    print(f"primes from {start} to {end}")
    i = 1
    for v in nums:
        print(f"{i}) {v}")
        i += 1

