--Quick sort
quickSort :: (Ord a) => [a] -> [a]
quickSort [] = []
quickSort (x:xs)=
  let smallerSort = quickSort [a | a <- xs, a <= x]
      biggerSort = quickSort [a | a <- xs, a > x]
  in smallerSort ++ [x] ++ biggerSort

--exec on command shell
--fmap (+1) [0..9]
