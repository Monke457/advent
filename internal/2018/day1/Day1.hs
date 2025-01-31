module Day1(main) where

main :: IO ()
main = do
  solveFirst "../../data/2018/day1.txt" 
 
  -- very very slow
  solveSecond "../../data/2018/day1.txt"


solveFirst filename = do
  contents <- readFile filename
  let linesOfFile = lines contents
      result = foldl processLine 0 linesOfFile
  putStrLn $ "First: " ++ show result 


solveSecond filename = do
  contents <- readFile filename
  let linesOfFile = lines contents
  let (result, _) = processUntilSeen (0, []) linesOfFile
  putStrLn $ "Second: " ++ show result 


processLine acc (op:rest) =
  case op of
    '+' -> acc + num
    '-' -> acc - num
    _ -> acc
  where
    num = read rest :: Int


processWithCheck (acc, seen, done) [] = (acc, seen, False) 
processWithCheck (acc, seen, done) (line:lines) =
  if acc `elem` seen 
  then (acc, seen, True)
  else processWithCheck (res, seen ++ [acc], False) lines
  where 
    res = processLine acc line


processUntilSeen (state, seen) linesOfFile = 
  let (res, allSeen, done) = processWithCheck (state, seen, done) linesOfFile 
  in if done
     then (res, allSeen)
     else processUntilSeen (res, allSeen) linesOfFile 
