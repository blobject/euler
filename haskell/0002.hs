-- file: 0002.hs
-- by  : agaric
-- desc: project euler #2 - "even Fibonacci numbers"

module Main (main) where

main :: IO ()
main = print answer

answer :: Int
answer = sum [x | x <- takeWhile (<= 4000000) fibs, even x]

fibs :: [Int]
fibs = 1 : 1 : zipWith (+) fibs (tail fibs)
