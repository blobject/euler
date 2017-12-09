/* file: 0004.rs
 * by  : agaric
 * copy: public domain
 * desc: project euler #4 - "largest palindrome product"
 * lang: rust
 */

fn is_palindrome(n: usize) -> bool {
    let mut v = Vec::new();
    {
        let mut a = n;
        while a >= 1 {
            v.push(a % 10);
            a /= 10;
        }
    }
    let mut l = Vec::new();
    let mut r = Vec::new();
    for i in 0..(v.len()) {
        if i < (v.len() / 2) {
            l.push(v[i]);
        } else {
            r.push(v[i]);
        }
    }
    if r.len() > l.len() {
        r.remove(0);
    }
    r.reverse();
    l.iter().zip(r).filter(|&(a, b)| *a != b).count() == 0
}

fn main() {
    let mut v = Vec::new();
    // presuming the answer lies somewhere 800..999
    for i in (800..1000).rev() {
        for j in (800..1000).rev() {
            if is_palindrome(i * j) {
                v.push([i * j, i, j]);
            }
        }
    }
    println!("{:?}", v.iter().max().unwrap());
}
