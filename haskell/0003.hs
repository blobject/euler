-- file: 0003.hs
-- by  : agaric
-- desc: project euler #3 - "largest prime factor"

module Main (main) where

main :: IO ()
main = print answer

answer :: Integer
answer = last (factors 600851475143)

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
