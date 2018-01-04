-- file: 0004.hs
-- by  : agaric
-- copy: public domain
-- desc: project euler #4 - "largest palindrome product"
-- lang: haskell

module Main (main) where

main :: IO ()
main = print answer

answer :: (Integer, Integer, Integer)
answer = maximum [(i * j, i, j)
                 | i <- [999,998..800], j <- [999,998..800],
                   isPalindrome (i * j)]

isPalindrome :: Integer -> Bool
isPalindrome n = s == reverse s where s = show n
