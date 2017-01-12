module DNA (toRNA) where

toRNA :: String -> Maybe String
toRNA = sequenceA . map translate
  where
    translate 'G' = Just 'C'
    translate 'C' = Just 'G'
    translate 'T' = Just 'A'
    translate 'A' = Just 'U'
    translate _   = Nothing
