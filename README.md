# Messari Market Data Coding Challenge - Solution


## run

see makefile

- run: run main programm with big input to evalueate perfomance
- test: run unit tests
- integration_test: run main programs with small input, to evaluate correcteness


## Decisions and explanations


I chose to follow the "official" project layout https://github.com/golang-standards/project-layout

For the internall folder i used a mix approach of clean architecture (application and usecase) and DDD (domain/entity)

I did not use any local database or queue because i understood this exercice as more like a "local script" aproach, and did not segrate any repository model for such

For the most performace part:
- i store the market summary after each trade is processed, avoiding to store all trades and the calculate all summaries
- i used a map to only need to store O(number of markets) of items
- after that i got a 20 seconds execution time on average, on my machine
- i introduce some goroutines to segretate the reading from input to process the input
- after this i got 15 seconds execution time on average, on my machine
- One improvement, i would think, would be to segregate per market te calculations, but this would require: 
  - a channel per market to receive from input
  - a channel per market to return the summary
  - a final iteration to "notify" each "worker" that the input has ENDED
  - after carefully examination, i would say this would requer a lot more code complexity for a small improvement on performance
  - hence the duration was only half a day, i am not pursuing tha optimal solution in my evaluation

