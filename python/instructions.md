# Getting started
Requirements: 
- pip
- python

## Cloning this repo
install git on your machine OR download this as a zip-archive and unpack it into your desired destination

if you installed git, move to your desired project directory and clone this Repo running 
    
    git clone https://github.com/DimitriFankhauser/Advent_Of_Dimitri

## creating a virtual enviroment and activating it
Move to the project-directory (Advent_Of_Dimitri) and then create a python virtual environment  by running this command: 

    python3 -m venv .venv

(Windows) activate the virtual environment by running the following script from powershell
    
    .venv\Scripts\Activate.ps1

When using an IDE like Jetbrains' Pycharm, you may still need to configure your python interpreter (.venv\Scripts\python.exe) through the UI, so you can run the tests from there

install all dependencies with: 

    pip install -r requirements.txt 

Now run  [exercise0](exercises/exercise0/ex0.py). This should print you the currently installed version of numpy and prove that everything is installed correctly

## Checking your solution 
There are unit-tests included in this repository. You can identify the test-files by this naming convention: 

    sol<Number>.py
You may run these test-files from the console (python3 sol<Number>.py) or through the user Interface in the IDE of your choice