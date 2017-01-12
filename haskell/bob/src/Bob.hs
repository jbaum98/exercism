module Bob (responseFor) where

import Data.Char

responseFor :: String -> String
responseFor utterance
  | isSilence = "Fine. Be that way!"
  | isYelling = "Whoa, chill out!"
  | isQuestion = "Sure."
  | otherwise = "Whatever."
  where
    isSilence  = all isSpace utterance
    isQuestion = last nonSpace == '?'
    isYelling  = not (null letters) && all isUpper letters
    nonSpace   = filter (not . isSpace) utterance
    letters    = filter isLetter utterance
