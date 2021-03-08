-- file: 0007.hs
-- by  : agaric
-- desc: project euler #7 - "10001st prime"

module Main (main) where

main :: IO ()
main = print answer

answer :: Integer
answer = primes !! 10000

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
