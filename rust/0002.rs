/* file: 0002.rs
 * by  : agaric
 * copy: public domain
 * desc: project euler #2 - "even Fibonacci numbers"
 * lang: rust
 */

fn fib(n: usize) -> Vec<usize> {
    let mut v = vec![1usize, 1];
    for i in 2..(n + 1) {
        let sum = v[i - 1] + v[i - 2];
        v.push(sum);
    }
    v
}

fn main() {
    let f = fib(32); // fib(33) > 4 mil
    println!("{}", f.iter()
             .filter(|&&x| x % 2 == 0)
             .fold(0, |sum, x| sum + x));
}
