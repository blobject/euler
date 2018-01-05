-- file: 0005.hs
-- by  : agaric
-- copy: public domain
-- desc: project euler #5 - "smallest multiple"
-- lang: haskell

module Main (main) where

main :: IO ()
main = print answer

answer :: Integer
answer = lcm' [1..20]

lcm' :: [Integer] -> Integer
lcm' [] = 0
lcm' (x:xs)
  | length xs > 1 = lcm x (lcm' xs)
  | otherwise     = lcm x (head xs)
