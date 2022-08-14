# Fasttrack Code Test

The idea of this project is developing a simple quiz using Golang and a Cobra CLI that talks with the API implemented by Francisco Augusto Cesar de Camargo Bellaz Guiraldelli.

## User stories/Use cases

The main demanded user stories and use cases for this project is described below:

    - User should be presented questions with a number of answers;
    
    - User should be able to select just one answer per question;
    
    - User should be able to answer all the questions and then post his/her answers and get back how many correct answers there had and be displayed to the user;
    
    - User should see how good he/she rated compared to others that had taken the quiz, e.g. "You scored higher than 60% of all quizzers".

## Aditional user stories / use cases

Additionally to make more easier the user experience and the development for this project was added also the user stories / use cases below:
    - User should be able to register your user using name and e-mail;

## Running the code
To run the project you have to follow the steps below:

1. Run the command to build compile the executable file in the main directory:

    ```console
    go build .
    ```
2. Run the command to start server:

    ```console
    <executable_filename> startServer
    ```
3. Run the command to start quiz:

    ```console
    <executable_filename> startQuiz
    ```

**ENJOY!**
