
if (process.argv.length < 4) {
  console.log("expected two arguments Example: node . 2 100 ")
}

const start = parseInt(process.argv[2])
const end = parseInt(process.argv[3])

const primes = [];

for (let n = start; n <= end; n++) {
  const root = Math.sqrt(n)
  let isPrime = true
  for (let i = 2; i <= root; i++) {
    if (n % i == 0) {
      isPrime = false;
      break
    }
  }
  if (isPrime) {
    primes.push(n)
  }
}

console.log(`${primes.length}) primes found from ${start} to ${end}`)
for (let i = 0; i < primes.length; i++) {
  console.log(`${i + 1}) ${primes[i]}`)
}




