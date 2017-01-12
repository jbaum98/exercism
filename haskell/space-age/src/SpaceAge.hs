module SpaceAge (Planet(..), ageOn) where

data Planet = Earth | Mercury | Venus | Mars | Jupiter | Saturn | Uranus | Neptune

yearLength :: Planet -> Float
yearLength Earth   = 31557600
yearLength Mercury = 0.2408467  * yearLength Earth
yearLength Venus   = 0.61519726 * yearLength Earth
yearLength Mars    = 1.8808158  * yearLength Earth
yearLength Jupiter = 11.862615  * yearLength Earth
yearLength Saturn  = 29.447498  * yearLength Earth
yearLength Uranus  = 84.016846  * yearLength Earth
yearLength Neptune = 164.79132  * yearLength Earth

ageOn :: Planet -> Float -> Float
ageOn = flip (/) . yearLength
