-- file: 0001.hs
-- by  : agaric
-- copy: public domain
-- desc: project euler #1 - "multiples of 3 and 5"
-- lang: haskell

module Main (main) where

main :: IO ()
main = print answer

answer :: Int
answer = sum [ x | x <- [3..999], rem x 3 == 0 || rem x 5 == 0 ]
