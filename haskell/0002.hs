-- file: 0002.hs
-- by  : agaric
-- copy: public domain
-- desc: project euler #2 - "even Fibonacci numbers"
-- lang: haskell

module Main (main) where

main :: IO ()
main = print answer

fibs :: [Int]
fibs = 1 : 1 : zipWith (+) fibs (tail fibs)

answer :: Int
answer = sum [ x | x <- takeWhile (<= 4000000) fibs, even x ]
