module Day2(main) where

main = do
  solveFirst "../../data/2018/day2.txt"

solveFirst filename = do
  contents <- readFile filename
  
  let linesOfFile = lines contents
  let (dub, trip) = foldl processLine (0, 0) linesOfFile
  putStrLn ("First: " ++ show (dub * trip))


processLine :: (Int, Int) -> String -> (Int, Int)
processLine (d, t) line = 
  let (trip, tripChar) = containsTriple line 
  in let (dub, dubChar) = containsDouble line
    in if dub && trip && tripChar == dubChar 
      then (d, t + 1)
      else if dub && trip 
      then (d + 1, t + 1) 
      else if trip 
      then (d, t + 1) 
      else if dub 
      then (d + 1, t)
      else (d, t)


containsDouble line = 
  -- todo: implement function
  (True, 'a')

containsTriple line = 
  -- todo: implement function
  (True, 'b')
