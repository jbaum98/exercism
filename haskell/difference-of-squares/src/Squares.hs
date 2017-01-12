{-# OPTIONS_GHC -fno-warn-type-defaults  #-}

module Squares
  (difference
  ,squareOfSums
  ,sumOfSquares)
  where

difference :: Integral a => a -> a
difference n = squareOfSums n - sumOfSquares n

squareOfSums :: Integral a => a -> a
squareOfSums n = (^ 2) $ n * (n + 1) `div` 2

sumOfSquares :: Integral a => a -> a
sumOfSquares n = n * (n + 1) * (2 * n + 1) `div` 6
