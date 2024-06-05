# Statistical Population Analysis

This Go program is designed to analyze a statistical population represented by a text file, where each line contains one value.

## Usage
### Prerequisites

   * Go installed on your system

### Steps to Run

1. Clone this repository to your local machine.

    ```
    git clone https://learn.zone01kisumu.ke/git/fnamayi/math-skills.git
    ```

2. Navigate to the directory containing the program.

```
cd math-skills
```

3. Run the program by providing a data file as input.
```
go run . data.txt
```

Replace **data.txt** with the path to your data file. The data file should contain **numerical data**, with each number separated by a new line or a comma.

### Note 
If the data file is empty or cannot be read, the program will display an error message and terminate.



## Input File Format  (data.txt)

The input file should follow these guidelines:

* **One Value Per Line:**   Each line in the file should contain one value representing a data point of the statistical population.

* **Numeric Values Only:**   Ensure that each value in the file is a numeric value. Non-numeric values will cause errors in the program.

* **Empty Lines:**   The program ignores empty lines in the file. You can include empty lines for readability or to separate sections within the file.


## Example Input File Format

Here's an example of how your input file (data.txt) might look:

```
23
45
67
12
8
1,000
40
```

## Output

After reading the file, the program executes several calculations and prints the results in the following format:

``` 
Average: 171
Median: 40
Variance: 114978
Standard Deviation: 339
```

Please note that the values are rounded integers.

## Additional Information

   * This program supports numerical data within the range of a 64-bit float.
   * Ensure that your data file contains valid numerical data for accurate analysis.
   * For any issues or feedback, please open an issue in the GitHub repository [NAMAYI](https://github.com/fnamayi).


