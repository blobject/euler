fn main() {
    println!("{}", (1..1000)
             .filter(|i| i % 3 == 0 || i % 5 == 0)
             .fold(0, |acc, x| acc + x));
}
