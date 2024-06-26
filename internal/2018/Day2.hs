module Day2(main) where
import Debug.Trace (trace)
import Data.Map (Map)
import qualified Data.Map as Map

main = do
  solveFirst "../../data/2018/day2.txt"

solveFirst filename = do
  contents <- readFile filename
  
  let linesOfFile = lines contents
  let (dub, trip) = foldl processLine (0, 0) linesOfFile
  putStrLn ("First: " ++ show (dub * trip))


processLine :: (Int, Int) -> String -> (Int, Int)
processLine (d, t) line =
  let counts = countLetters line
  in let trip = containsTriple counts
         dub = containsDouble counts 
    in if trip && dub
      then (d + 1, t + 1) 
      else if trip 
      then (d, t + 1) 
      else if dub 
      then (d + 1, t)
      else (d, t)


containsDouble :: Map Char Int -> Bool
containsDouble counts = 2 `elem` Map.elems counts


containsTriple :: Map Char Int -> Bool
containsTriple counts = 3 `elem` Map.elems counts


countLetters :: String -> Map Char Int
countLetters = foldr (\c m -> Map.insertWith (+) c 1 m) Map.empty
