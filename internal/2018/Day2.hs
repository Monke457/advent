module Day2(main) where
import Debug.Trace (trace)

main = do
  solveFirst "../../data/2018/day2.txt"

solveFirst filename = do
  contents <- readFile filename
  
  let linesOfFile = lines contents
  let (dub, trip) = foldl processLine (0, 0) linesOfFile
  putStrLn ("First: " ++ show (dub * trip))


processLine :: (Int, Int) -> String -> (Int, Int)
processLine (d, t) line = 
  let (trip, cTrip) = containsTriple line 
  in let (dub, cDub) = containsDouble line
    in trace ("Processing line: " ++ line ++ " trip info: " ++ show trip ++ " " ++ show cTrip ++ " " ++ show t ++ " dub info: " ++ show dub ++ " " ++ show cDub ++ " " ++ show d) $ 
      if dub && trip && cTrip == cDub
      then (d, t + 1)
      else if dub && trip 
      then (d + 1, t + 1) 
      else if trip 
      then (d, t + 1) 
      else if dub 
      then (d + 1, t)
      else (d, t)


containsDouble line = 
  (True, 'a')

containsTriple line = 
  (True, 'b')
