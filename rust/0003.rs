/* file: 0003.rs
 * by  : agaric
 * desc: project euler #3 - "largest prime factor"
 */

fn sieve(n: usize) -> Vec<usize> {
    let mut v = vec![true; n + 1];
    v[0] = false;
    if n >= 1 { v[1] = false; }

    for i in 2..(n + 1) {
        if v[i] {
            let mut m = i * i;
            while m <= n {
                v[m] = false;
                m += i;
            }
        }
    }

    v.iter().enumerate()
        .filter_map(|(p, &is_p)| if is_p {Some(p)} else {None})
        .collect()
}

fn main() {
    let x = 600851475143usize;
    let primes = sieve((x as f64).sqrt() as usize);
    let mut max = 1;
    for p in primes {
        if x % p == 0 && p > max {
            max = p;
        }
    }
    println!("{}", max);
}
