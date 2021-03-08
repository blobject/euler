/* file: 0001.rs
 * by  : agaric
 * desc: project euler #1 - "multiples of 3 and 5"
 */

fn main() {
    println!("{}", (1..1000)
             .filter(|i| i % 3 == 0 || i % 5 == 0)
             .fold(0, |sum, x| sum + x));
}
