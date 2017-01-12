module SumOfMultiples (sumOfMultiples) where

sumOfMultiples :: (Foldable f, Integral a) => f a -> a -> a
sumOfMultiples factors lim =
  sum [x | x <- [1..pred lim], any (divBy x) factors]
  where
    n `divBy` f = n `mod` f == 0
