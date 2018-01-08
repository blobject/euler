-- file: 0010.hs
-- by  : agaric
-- copy: public domain
-- desc: project euler #10 - "summation of primes"
-- lang: haskell

module Main (main) where

main :: IO ()
main = print answer

answer :: Integer
answer = sum (takeWhile (< 2000000) primes)

primes :: [Integer]
primes = 2 : filter (null . tail . factors) [3, 5..]

factors :: Integer -> [Integer]
factors n = factor n primes

factor :: Integer -> [Integer] -> [Integer]
factor _ [] = []
factor n (p:ps)
  | p * p > n    = [n]
  | rem n p == 0 = p : factor (div n p) (p:ps)
  | otherwise    = factor n ps
