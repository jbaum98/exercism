module Grains (square, total) where

import Data.Maybe

square :: Integral a => a -> Maybe a
square n
  | n <= 0 || n > 64 = Nothing
  | otherwise      = Just $ 2^(n-1)

total :: Integral a => a
total = sum $ (fromMaybe 0 . square) <$> [1..64]
