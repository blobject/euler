-- file: 0009.hs
-- by  : agaric
-- copy: public domain
-- desc: project euler #9 - "special pythagorean triplet"
-- lang: haskell

module Main (main) where

main :: IO ()
main = print answer

answer :: Int
answer = product (head (filter (\[a, b, c] -> a + b + c == 1000) triplets))

triplets :: [[Int]]
triplets = [[floor a, floor b, floor c]
           | a <- [1..1000]
           , b <- [1..a-1]
           , let c = sqrt (a^2 + b^2)
           , floor c == ceiling c]
