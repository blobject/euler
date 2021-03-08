-- file: 0001.hs
-- by  : agaric
-- desc: project euler #1 - "multiples of 3 and 5"

module Main (main) where

main :: IO ()
main = print answer

answer :: Int
answer = sum [x | x <- [3..999], rem x 3 == 0 || rem x 5 == 0]
